package models

import (
	"time"
	"labix.org/v2/mgo/bson"
)

type User struct {
	Id 		bson.ObjectId	`bson:"_id"`
	Name 		string		`bson:"name"`
	CreatedDate	time.Time	`bson:"createdDate"`
}

