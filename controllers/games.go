package controllers

import (
	"../services"
	"../models"
	"encoding/json"
)

func GetGames() ([]byte) {
	var games = []models.Game{}
	games = services.GetGames()

	for i, game := range games {
		games[i].Player1 = GetUserById(game.Player1Id)
		games[i].Player2 = GetUserById(game.Player2Id)
	}

	str, _ := json.Marshal(games)

	return str;
}
