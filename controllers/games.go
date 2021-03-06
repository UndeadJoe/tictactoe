package controllers

import (
	"encoding/json"
	"github.com/go-martini/martini"
	"labix.org/v2/mgo/bson"
	"net/http"
	"tictactoe/config"
	"tictactoe/models"
	"tictactoe/services"
	"tictactoe/utils"
)

type resultData struct {
	Id     bson.ObjectId `json:"_id"`
	Title  string        `json:"title"`
	Status int           `json:"status"`
}

func GetGames() (str []byte) {
	var games = []models.Game{}
	var data = []resultData{}
	games = services.GetGames()

	for _, game := range games {
		data = append(data, resultData{
			Id:     game.Id,
			Title:  game.Title,
			Status: game.Status})
	}
	result := map[string]interface{}{"status": "ok", "data": data}
	str, _ = json.Marshal(result)

	return
}

// population of player fields
// TODO: rewrite for native population method
func populatePlayers(game *models.Game) (err config.ApiError) {
	if game.Player1Id.Valid() {
		game.Player1, err = GetUserById(game.Player1Id)
	}

	if game.Player2Id.Valid() {
		game.Player2, err = GetUserById(game.Player2Id)
	}
	return
}

func getGameById(id bson.ObjectId) (game models.Game, err config.ApiError) {

	// TODO: rewrite to one return
	if !id.Valid() {
		err = config.ErrGameIdWrong
		return
	}

	game, err = services.GetGame(id)
	if err.Code != 0 {
		return
	}

	_ = populatePlayers(&game)
	return
}

func GetGame(params martini.Params) (str []byte) {
	var game = models.Game{}
	var err config.ApiError
	var id = bson.ObjectIdHex(params["id"])

	game, err = getGameById(id)

	if err.Code != 0 {
		result := map[string]interface{}{"status": "error", "error": err}
		str, _ = json.Marshal(result)
		return
	}

	result := map[string]interface{}{"status": "ok", "game": game}
	str, _ = json.Marshal(result)

	return
}

func CreateGame(res http.ResponseWriter, req *http.Request) (str []byte) {
	var (
		params      = utils.BodyToStruct(req)
		game        = models.Game{}
		user        = models.User{}
		username    = params["username"].(string)
		accessToken = req.Header.Get("x-token")
		result      = map[string]interface{}{}
	)

	user = FindUser(accessToken, username)
	game, err := game.Create(params, user)
	if err.Code != 0 {
		result = map[string]interface{}{"status": "error", "error": err}
	} else {
		game.Id, _ = services.AddGame(game)
		result = map[string]interface{}{"status": "ok", "game": game, "access_token": user.Id.Hex()}
	}
	str, _ = json.Marshal(result)
	return
}

func JoinGame(res http.ResponseWriter, req *http.Request, params martini.Params) (str []byte) {
	var (
		id          = bson.ObjectIdHex(params["id"])
		reqParams   = utils.BodyToStruct(req)
		username    = reqParams["username"]
		accessToken = req.Header.Get("x-token")
		result      map[string]interface{}
		user        = models.User{}
		game        = models.Game{}
		err         config.ApiError
	)

	if username == nil && accessToken == "" {
		result = map[string]interface{}{"status": "error", "error": config.ErrNoUser}
	} else {
		user = FindUser(accessToken, username.(string))

		game, err = getGameById(id)
		if err.Code != 0 {
			result = map[string]interface{}{"status": "error", "error": err}
		}

		game, err = services.JoinGame(id, user.Id)
		if err.Code != 0 {
			result = map[string]interface{}{"status": "error", "error": err}
		} else {
			_ = populatePlayers(&game)
			result = map[string]interface{}{"status": "ok", "game": game, "access_token": user.Id.Hex()}
		}
	}

	str, _ = json.Marshal(result)
	return
}

func MakeMove(res http.ResponseWriter, req *http.Request, params martini.Params) (str []byte) {
	var (
		result      map[string]interface{}
		reqParams   = utils.BodyToStruct(req)
		id          = bson.ObjectIdHex(params["id"])
		accessToken = req.Header.Get("x-token")
		row         = int(reqParams["row"].(float64))
		col         = int(reqParams["col"].(float64))
		game        = models.Game{}
		err         config.ApiError
		nextTurn    = 0
		userIndex   = 0
	)

	game, err = getGameById(id)
	if err.Code != 0 {
		result = map[string]interface{}{"status": "error", "error": err}
		str, _ = json.Marshal(result)
		return
	}

	if game.CheckActive() == false {
		result = map[string]interface{}{"status": "error", "error": config.ErrNoActiveGame}
		str, _ = json.Marshal(result)
		return
	}

	if FindUser(accessToken, "").Name == "" {
		result = map[string]interface{}{"status": "error", "error": config.ErrNoUser}
		str, _ = json.Marshal(result)
		return
	} else {
		switch accessToken {
		case game.Player1Id.Hex():
			userIndex = 1
			nextTurn = 2
			break
		case game.Player2Id.Hex():
			userIndex = 2
			nextTurn = 1
			break
		default:
			userIndex = 0
		}

		if userIndex == 0 {
			result = map[string]interface{}{"status": "error", "error": config.ErrBadPlayer}
			str, _ = json.Marshal(result)
			return
		}
	}

	if (row >= game.PoleSize || col >= game.PoleSize || game.Field[row][col].State != 0) {
		result = map[string]interface{}{"status": "error", "error": config.ErrBadCell}
		str, _ = json.Marshal(result)
		return
	}

	if game.CurrentTurn != userIndex {
		result = map[string]interface{}{"status": "error", "error": config.ErrBadPlayer}
		str, _ = json.Marshal(result)
		return
	}

	game.CurrentTurn = nextTurn
	r := services.MakeMove(&game, row, col, userIndex)
	if r == false {
		result = map[string]interface{}{"status": "error", "error": config.ErrBadTurn}
	} else {
		// TODO: Сделать опеределение победителя
		game.WinnerCheck(row, col)
		result = map[string]interface{}{"status": "ok", "field": game.Field,
			"winnerIndex": game.WinnerIndex, "winnerName": game.WinnerName}
	}

	str, _ = json.Marshal(result)
	return
}
