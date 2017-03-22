package controllers

import (
	"tictactoe/models"
	"tictactoe/services"
	"tictactoe/config"
	"encoding/json"
	"labix.org/v2/mgo/bson"
)

func GetUserById(id bson.ObjectId) (user models.User, err config.ApiError) {
	user, e := services.GetUser(id)
	if e != nil {
		err = config.ErrNoUser
	}
	return
}


func GetUsers() ([]byte) {
	var users = []models.User{}
	users = services.GetUsers()

	str, _ := json.Marshal(users)

	return str
}

func CreateUser(accessToken string, username string) (user models.User) {
	if accessToken == "" {
		accessToken = "000000000000000000000000"}

	user = services.AddUser(bson.ObjectIdHex(accessToken), username)
	return user
}