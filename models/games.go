package models

import (
	"labix.org/v2/mgo/bson"
	"time"
	"log"
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

func (p *Game) Create(data map[string]interface{}, user User) (game Game) {
	var (
		title = data["title"].(string)
		poleSize = data["poleSize"]
		poleSizeInt = 3
	)

	if poleSize != nil {
		poleSizeInt = int(poleSize.(float64))}

	game = Game{
		Title: title,
		PoleSize: poleSizeInt,
		Field: make([][]Field, poleSizeInt),
		Player1Id: user.Id,
		Player1: user,
	}

	// make field array
	for i := 0; i < poleSizeInt; i++ {
		game.Field[i] = make([]Field, poleSizeInt)
	}

	log.Println(game.Player1)

	return game
}