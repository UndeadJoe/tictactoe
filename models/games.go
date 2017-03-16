package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Game struct {
	Id		bson.ObjectId	`json:"_id" bson:"_id"`
	Title 		string		`json:"title" bson:"title"`
	Status 		int		`json:"status" bson:"status"`
	PoleSize 	int		`json:"poleSize" bson:"poleSize"`
	CurrentTurn	int		`json:"currentTurn" bson:"currentTurn"`

	WinnerIndex	int		`json:"winnerIndex" bson:"winnerIndex"`
	WinnerName	string		`json:"winnerName" bson:"winnerName"`

	Field		[][]Field	`json:"field" bson:"field"`

	CreatedDate	time.Time	`json:"createdDate" bson:"createdDate"`

	Player1Id	bson.ObjectId	`json:"player1" bson:"player1"`
	Player2Id	bson.ObjectId	`json:"player2" bson:"player2"`
	Player1		User		`json:"-" bson:"-"`
	Player2		User		`json:"-" bson:"-"`
}

type Field struct {
	State 		int	`json:"state" bson:"state"`
}