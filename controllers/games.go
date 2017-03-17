package controllers

import (
	"../services"
	"../models"
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"github.com/go-martini/martini"
	"log"
)

//{"status":"ok","data":[{"_id":"588af2478081d52dec3b79c5","title":"тест","status":20}]}

type resultData struct {
	Id	bson.ObjectId	`json:"_id"`
	Title	string		`json:"title"`
	Status	int		`json:"status"`
}

func GetGames() ([]byte) {
	var games = []models.Game{}
	var data = []resultData{}
	games = services.GetGames()

	for _, game := range games {
		data = append(data,  resultData {
			Id: game.Id,
			Title: game.Title,
			Status: game.Status})
	}
	result := map[string]interface{} {"status": "ok", "data": data}
	str, _ := json.Marshal(result)

	return str;
}

func GetGameById(params martini.Params) ([]byte) {
	var game = models.Game{}
	var id = bson.ObjectIdHex(params["id"])
	game = services.GetGame(id)

	// population of player fields
	// TODO: rewrite for native population method
	game.Player1 = GetUserById(game.Player1Id)
	game.Player2 = GetUserById(game.Player2Id)

	result := map[string]interface{} {"status": "ok", "game": game}
	str, _ := json.Marshal(result)

	return str;
}

func CreateGame(params martini.Params) {
	var (
		gameTitle = params["title"]
		userName = params["username"]
		poleSize = params["polesize"]
	)

	log.Println(gameTitle, userName, poleSize)
}