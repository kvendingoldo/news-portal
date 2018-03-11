package endpoints

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"

	dao "github.com/kvendingoldo/news-portal/dao"
	. "github.com/kvendingoldo/news-portal/models"
)

var Dao = dao.Init()

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func AllPostsEndPoint(w http.ResponseWriter, r *http.Request) {
	posts, err := Dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, posts)
}

func FindPostEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	post, err := Dao.FindById(params["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Post ID")
		return
	}
	respondWithJson(w, http.StatusOK, post)
}

func CreatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	post.ID = bson.NewObjectId()
	if err := Dao.Insert(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, post)
}

func UpdatePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Dao.Update(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}

func DeletePostEndPoint(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var post Post
	if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	if err := Dao.Delete(post); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, map[string]string{"result": "success"})
}
