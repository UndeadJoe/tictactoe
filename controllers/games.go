package controllers

import (
	"tictactoe/services"
	"tictactoe/models"
	"tictactoe/utils"
//	"../config"
	"encoding/json"
	"labix.org/v2/mgo/bson"
	"github.com/go-martini/martini"
	//"log"
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
	var err = ""
	var id = bson.ObjectIdHex(params["id"])
	// TODO: rewrite to one return
	if id.Valid() {
		game, err = services.GetGame(id)

		if err != "" {
			result := map[string]interface{} {"status": "error", "error": err}
			str, _ := json.Marshal(result)
			return str
		}
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


func CreateGame(res http.ResponseWriter, req *http.Request) (str []byte) {
	var (
		params = utils.BodyToStruct(req)
		game = models.Game{}
		user = models.User{}
		username = params["username"].(string)
		accessToken = req.Header.Get("x-token")
		err = ""
		result = map[string]interface{} {}
	)

	user = CreateUser(accessToken, username)
	game, err = game.Create(params, user)
	if err != "" {
		result = map[string]interface{} {"status": "error", "error": err}
	} else {
		// TODO: Доделать обработку ошибок
		game.Id, _ = services.AddGame(game)
		result = map[string]interface{} {"status": "ok", "game": game, "access_token": user.Id.Hex()}
	}
	str, _ = json.Marshal(result)

	return str
}