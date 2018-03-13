package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	. "github.com/kvendingoldo/news-portal/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/news", AllPostsEndPoint).Methods("GET")
	router.HandleFunc("/news", CreatePostEndPoint).Methods("POST")
	router.HandleFunc("/news", UpdatePostEndPoint).Methods("PUT")
	router.HandleFunc("/news", DeletePostEndPoint).Methods("DELETE")
	router.HandleFunc("/news/{id}", FindPostEndpoint).Methods("GET")
	if err := http.ListenAndServe(":8000", router); err != nil {
		log.Fatal(err)
	}
}
