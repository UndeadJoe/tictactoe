package models

import (
	"labix.org/v2/mgo/bson"
	"tictactoe/config"
	"time"
)

type Game struct {
	Id          bson.ObjectId `json:"_id" bson:"_id,omitempty"`
	Title       string        `json:"title" bson:"title"`
	Status      int           `json:"status" bson:"status"`
	PoleSize    int           `json:"poleSize" bson:"poleSize"`
	CurrentTurn int           `json:"currentTurn" bson:"currentTurn"`

	WinnerIndex int    `json:"winnerIndex" bson:"winnerIndex"`
	WinnerName  string `json:"winnerName" bson:"winnerName"`

	Field [][]Field `json:"field" bson:"field"`

	CreatedDate time.Time `json:"createdDate" bson:"createdDate"`

	Player1Id bson.ObjectId `json:"player1id" bson:"player1,omitempty"`
	Player2Id bson.ObjectId `json:"player2id" bson:"player2,omitempty"`
	Player1   User          `json:"player1" bson:"-"`
	Player2   User          `json:"player2" bson:"-"`
}

type Field struct {
	State int `json:"state" bson:"state"`
}

func GameStatus() map[string]interface{} {
	result := map[string]interface{}{
		"new":      0,
		"active":   10,
		"finished": 20,
		"deleted":  30}

	return result
}

func (p *Game) Create(data map[string]interface{}, user User) (Game, config.ApiError) {
	var (
		title       = data["title"]
		poleSize    = data["poleSize"]
		poleSizeInt = 3
		status      = GameStatus()["new"].(int)
	)

	if title == nil {
		return Game{}, config.ErrGameTitleWrong
	}

	if poleSize != nil {
		poleSizeInt = int(poleSize.(float64))
	}

	game := Game{
		Title:       title.(string),
		PoleSize:    poleSizeInt,
		Field:       make([][]Field, poleSizeInt),
		Player1Id:   user.Id,
		Player1:     user,
		Status:      status,
		CreatedDate: time.Now(),
		CurrentTurn: 1}

	// make field array
	for i := 0; i < poleSizeInt; i++ {
		game.Field[i] = make([]Field, poleSizeInt)
	}

	return game, config.ApiError{}
}

func (p *Game) JoinGame(id bson.ObjectId) (err config.ApiError) {
	p.Player2Id = id
	return
}

func (p *Game) CheckActive() bool {
	if p.Status == GameStatus()["active"] {
		return true
	}
	return false
}

// Winner check after players turn
func (p *Game) WinnerCheck(row int, col int) (winnerIndex int) {
	var (
		rowSum = 0
		colSum = 0
		diag1Sum = 0
		diag2Sum = 0
		poleSize = len(p.Field[row])
		playerIndex = p.Field[row][col].State
	)

	for i:=0; i < poleSize; i++ {
		if p.Field[row][i].State == playerIndex {
			rowSum++
		}
		if p.Field[i][col].State == playerIndex {
			colSum++
		}
		if p.Field[i][i].State == playerIndex {
			diag1Sum++
		}
		if p.Field[i][poleSize-i-1].State == playerIndex {
			diag2Sum++
		}
	}
	if (rowSum == poleSize) || (colSum == poleSize) || (diag1Sum == poleSize) || (diag2Sum == poleSize)  {
		winnerIndex = playerIndex
	}

	return
}