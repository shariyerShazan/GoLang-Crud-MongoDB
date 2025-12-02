package models

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID     bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string        `json:"name" bson:"name"`
	Email  string        `json:"email" bson:"email"`
	Age    int           `json:"age" bson:"age"`
	Gender string        `json:"gender" bson:"gender"`
}
