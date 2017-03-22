package services

import (
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
	"tictactoe/models"
	"tictactoe/config"
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

func GetGame(id bson.ObjectId) (models.Game, string) {
	connection := session.DB("tictactoe").C("games")
	result := models.Game{}
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		result = models.Game{}
		return result, config.ErrGameIdWrong.Error()
	}

	return result, "";
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
	// insert new user or update current
	_, err := connection.Upsert(
		bson.M{"$or": []bson.M{ bson.M{"_id": id}, bson.M{"name": username} }},
		bson.M{"name": username})
	if err != nil {
		log.Fatal(err)
	}

	err = connection.Find(bson.M{"name": username}).One(&result)
	if err != nil {
		log.Fatal(err)
	}

	return result
}

func AddGame(game models.Game) (newId bson.ObjectId, err config.ApiError) {
	connection := session.DB("tictactoe").C("games")
	info, e := connection.Upsert(bson.M{"title": ""}, game)
	if info.Updated == 0 {
		err = config.ErrCreateGame
	}
	newId = info.UpsertedId.(bson.ObjectId)
	if e != nil {
		log.Fatal(e)
	}

	return
}