package dao

import (
	"log"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	. "github.com/kvendingoldo/news-portal/config"
	. "github.com/kvendingoldo/news-portal/models"
)

type PostsDAO struct {
	Server     string
	Database   string
	Collection string
}

var db *mgo.Database

func Init() PostsDAO {
	var dao = PostsDAO{}
	var config = Config{}
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Collection = config.Collection
	dao.Connect()
	return dao
}

func (m *PostsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *PostsDAO) FindAll() ([]Post, error) {
	var posts []Post
	err := db.C(m.Collection).Find(bson.M{}).All(&posts)
	return posts, err
}

func (m *PostsDAO) FindById(id string) (Post, error) {
	var post Post
	err := db.C(m.Collection).FindId(bson.ObjectIdHex(id)).One(&post)
	return post, err
}

func (m *PostsDAO) Insert(post Post) error {
	err := db.C(m.Collection).Insert(&post)
	return err
}

func (m *PostsDAO) Delete(post Post) error {
	err := db.C(m.Collection).Remove(&post)
	return err
}

func (m *PostsDAO) Update(post Post) error {
	err := db.C(m.Collection).UpdateId(post.ID, &post)
	return err
}
