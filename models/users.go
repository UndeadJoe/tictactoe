package models

import (
	"time"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id 		bson.ObjectId	`json:"_id" bson:"_id"`
	Name 		string		`json:"name" bson:"name"`
	CreatedDate	time.Time	`json:"createdDate" bson:"createdDate"`
}

