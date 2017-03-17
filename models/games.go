package models

import (
	"labix.org/v2/mgo/bson"
	"time"
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

