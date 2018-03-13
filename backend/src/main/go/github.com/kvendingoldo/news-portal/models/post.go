package post

import (
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID          	bson.ObjectId `bson:"_id" json:"id"`
	Date 					string 				`bson:"date" json:"date"`
	Title       	string        `bson:"title" json:"title"`
	Author      	string        `bson:"author" json:"author"`
	Category	  	string 				`bson:"category" json:"category"`
	Announcement	string        `bson:"announcement" json:"announcement"`
	Body        	string        `bson:"body" json:"body"`
}
