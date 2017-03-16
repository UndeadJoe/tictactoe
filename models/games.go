package models

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type Game struct {
	Id		bson.ObjectId	`bson:"_id"`
	Title 		string	`bson:"title"`
	Status 		int	`bson:"status"`
	PoleSize 	int	`bson:"poleSize"`
	CurrentTurn	int	`bson:"currentTurn"`

	WinnerIndex	int	`bson:"winnerIndex"`
	WinnerName	string	`bson:"winnerName"`

	Field		[][]int	`bson:"field"`

	CreatedDate	time.Time	`bson:"createdDate"`

	Player1Id	bson.ObjectId	`bson:"player1"`
	Player2Id	bson.ObjectId	`bson:"player2"`
	Player1		User	`bson:"-"`
	Player2		User	`bson:"-"`
}