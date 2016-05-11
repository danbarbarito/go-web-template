package models

import "gopkg.in/mgo.v2/bson"

type (
	User struct {
		Name   string `json:"name" bson:"name"`
		Id     bson.ObjectId `json:"id" bson:"_id"`
	}
)
