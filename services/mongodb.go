package services

import (
	"tictactoe/models"
	"tictactoe/config"
	"labix.org/v2/mgo"
	"labix.org/v2/mgo/bson"
	"log"
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

func GetGame(id bson.ObjectId) (models.Game, config.ApiError) {
	connection := session.DB("tictactoe").C("games")
	result := models.Game{}
	err = connection.Find(bson.M{"_id" : id}).One(&result)
	if err != nil {
		result = models.Game{}
		return result, config.ErrGameIdWrong
	}

	return result, config.ApiError{}
}

func GetUser(id bson.ObjectId) (result models.User, err error) {
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{"_id" : id}).One(&result)

	return
}

func GetUsers() (result []models.User) {
	connection := session.DB("tictactoe").C("users")
	err = connection.Find(bson.M{}).All(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func AddUser(username string) (result models.User, err config.ApiError) {
	connection := session.DB("tictactoe").C("users")
	// insert new user or update current
	info, e := connection.Upsert(bson.M{"username": ""}, bson.M{"name": username})
	if e != nil {
		log.Fatal(e)
	}

	if (info.UpsertedId == nil) {
		err = config.ErrCreateUser
	} else {
		result = models.User{}
		result = result.Create(username)
		result.Id = info.UpsertedId.(bson.ObjectId)
	}

	return
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

func JoinGame(gameId bson.ObjectId, userId bson.ObjectId) (game models.Game, err config.ApiError) {
	connection := session.DB("tictactoe").C("games")
	change := mgo.Change{
		Update: bson.M{"$set": bson.M{"player2": userId}},
		ReturnNew: true}

	info, e := connection.Find(bson.M{"_id": gameId}).Apply(change, &game)

	log.Println(game)
	if info.Updated == 0 {
		err = config.NewApiError(e)
	}

	return
}