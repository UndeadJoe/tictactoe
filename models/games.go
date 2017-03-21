package models

import (
	"labix.org/v2/mgo/bson"
	"time"
	"log"
	"../config"
)

type Game struct {
	Id		bson.ObjectId	`json:"_id" bson:"_id,omitempty"`
	Title 		string		`json:"title" bson:"title"`
	Status 		int		`json:"status" bson:"status"`
	PoleSize 	int		`json:"poleSize" bson:"poleSize"`
	CurrentTurn	int		`json:"currentTurn" bson:"currentTurn"`

	WinnerIndex	int		`json:"winnerIndex" bson:"winnerIndex"`
	WinnerName	string		`json:"winnerName" bson:"winnerName"`

	Field		[][]Field	`json:"field" bson:"field"`

	CreatedDate	time.Time	`json:"createdDate" bson:"createdDate"`

	Player1Id	bson.ObjectId	`json:"player1id" bson:"player1,omitempty"`
	Player2Id	bson.ObjectId	`json:"player2id" bson:"player2,omitempty"`
	Player1		User		`json:"player1" bson:"-"`
	Player2		User		`json:"player2" bson:"-"`
}

type Field struct {
	State 		int	`json:"state" bson:"state"`
}

func GameStatus() map[string]interface{} {
	result := map[string]interface{} {
		"new": 0,
		"active": 10,
		"finished": 20,
		"deleted": 30}

	return result
}

func (p *Game) Create(data map[string]interface{}, user User) (Game, string) {
	var (
		title = data["title"]
		poleSize = data["poleSize"]
		poleSizeInt = 3
		status = GameStatus()["new"].(int)
	)

	if title == nil {
		return Game{}, config.ErrGameTitleWrong.Error()}

	if poleSize != nil {
		poleSizeInt = int(poleSize.(float64))}

	game := Game{
		Title: title.(string),
		PoleSize: poleSizeInt,
		Field: make([][]Field, poleSizeInt),
		Player1Id: user.Id,
		Player1: user,
		Status: status}

	// make field array
	for i := 0; i < poleSizeInt; i++ {
		game.Field[i] = make([]Field, poleSizeInt)
	}

	log.Println(game.Player1)

	return game, ""
}