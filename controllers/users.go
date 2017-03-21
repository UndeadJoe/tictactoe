package controllers

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"../models"
	"../services"
)

func GetUserById(id bson.ObjectId) (user models.User) {
	user = services.GetUser(id)

	//str, _ := json.Marshal(user)

	return user;
}


func GetUsers() ([]byte) {
	var users = []models.User{}
	users = services.GetUsers()

	str, _ := json.Marshal(users)

	return str;
}

func CreateUser(accessToken string, username string) (user models.User) {
	if accessToken == "" {
		accessToken = "000000000000000000000000"}

	user = services.AddUser(bson.ObjectIdHex(accessToken), username)
	return user
}