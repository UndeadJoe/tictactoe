package controllers

import (
	"../services"
	"../models"
	"../utils"
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"github.com/go-martini/martini"
	"log"
	"net/http"
)

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
	if id.Valid() {
		game = services.GetGame(id)
	}

	// population of player fields
	// TODO: rewrite for native population method
	game.Player1 = GetUserById(game.Player1Id)
	if game.Player2Id.Valid() {
		game.Player2 = GetUserById(game.Player2Id)
	}

	result := map[string]interface{} {"status": "ok", "game": game}
	str, _ := json.Marshal(result)

	return str;
}


func CreateGame(res http.ResponseWriter, req *http.Request) {
	type newGame struct {
		Title		string
		Username	string
		AccessToken	string
		PoleSize	int
		Field		[][]models.Field
	}

	result := utils.BodyToStruct(req)
	access_token := result["access_token"]
	poleSize := result["poleSize"]
	if access_token == nil {
		access_token = bson.NewObjectId().String()}
	if poleSize == nil {
		poleSize = 3.0}

	// new game initialize
	game := newGame{
		Title: result["title"].(string),
		Username: result["username"].(string),
		AccessToken: access_token.(string),
		PoleSize: int(poleSize.(float64)),
		Field: make([][]models.Field, int(poleSize.(float64)))}

	// make field array
	for i := 0; i < int(poleSize.(float64)); i++ {
		game.Field[i] = make([]models.Field, int(poleSize.(float64)))
	}

	log.Println(game)
}