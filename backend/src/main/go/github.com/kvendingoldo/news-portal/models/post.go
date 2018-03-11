package post

import (
	"gopkg.in/mgo.v2/bson"
)

type Post struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	CoverImage  string        `bson:"cover_image" json:"cover_image"`
	Description string        `bson:"description" json:"description"`
	Body        string        `bson:"body" json:"body"`
}
