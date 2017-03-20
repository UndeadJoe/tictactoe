package services

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"../models"
	"time"
)

var (
	session, err = mgo.Dial("localhost")

)

func init() {
	if err != nil {
		log.Fatal(err)
	}
}

func GetGames() (result []models.Game) {
	connection := session.DB("tictactoe").C("games")
	err = connection.Find(bson.M{}).All(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result;
}

func GetGame(id bson.ObjectId) (result models.Game) {
	connection := session.DB("tictactoe").C("games")
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		log.Println(err)
		result = models.Game{}
	}

	return result;
}

func GetUser(id bson.ObjectId) (result models.User) {
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result;
}

func GetUsers() (result []models.User) {
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{}).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result;
}

func AddUser(id bson.ObjectId, username string) (result models.User) {
	connection := session.DB("tictactoe").C("users")
	users, _ := connection.Find(bson.M{	"$or": []bson.M{
		bson.M{"name": username},
		bson.M{"_id": id} } } ).Count()

	if users == 0 {
		err = connection.Insert(bson.M{"name": username, "createdDate": time.Now()})
	}
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Find(bson.M{"name": username}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}