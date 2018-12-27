package model

import (
	"gopkg.in/mgo.v2/bson"
)

type User struct {
	Id_ bson.ObjectId `bson:"_id"`
	Name string `bson:"name"`
	Age int `bson:"age"`
}

