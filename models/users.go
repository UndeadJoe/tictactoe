package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id          bson.ObjectId `json:"_id" bson:"_id"`
	Name        string        `json:"name" bson:"name"`
	CreatedDate time.Time     `json:"createdDate" bson:"createdDate"`
}

func (p *User) Create(username string) (user User) {

	user = User{
		Id:          bson.NewObjectId(),
		Name:        username,
		CreatedDate: time.Now()}

	return user
}
