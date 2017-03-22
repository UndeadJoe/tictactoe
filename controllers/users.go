package controllers

import (
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"tictactoe/config"
	"tictactoe/models"
	"tictactoe/services"
)

func GetUserById(id bson.ObjectId) (user models.User, err config.ApiError) {
	user, e := services.GetUser(id)
	if e != nil {
		err = config.ErrNoUser
	}
	return
}

func GetUsers() []byte {
	var users = []models.User{}
	users = services.GetUsers()

	str, _ := json.Marshal(users)

	return str
}

func FindUser(accessToken string, username string) (user models.User) {
	if accessToken == "" {
		user, _ = services.AddUser(username)
	} else {
		user, _ = services.GetUser(bson.ObjectIdHex(accessToken))
	}

	return
}
