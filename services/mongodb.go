package services

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"../models"
)

var (
	session, err = mgo.Dial("localhost")

)

func init() {
	if err != nil {
		log.Fatal(err)
	}
}

func GetGames() ([]models.Game) {
	result := []models.Game{}
	connection := session.DB("tictactoe").C("games")
	err = connection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result;
}

func GetGame(id bson.ObjectId) (models.Game) {
	result := models.Game{}
	connection := session.DB("tictactoe").C("games")
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		log.Println(err)
		result = models.Game{}
	}

	return result;
}

func GetUser(id bson.ObjectId) (models.User) {
	result := models.User{}
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result;
}

func GetUsers() ([]models.User) {
	result := []models.User{}
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{}).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result;
}