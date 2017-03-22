package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"net/http"
	"os"
	"tictactoe/controllers"
	"time"
)

func main() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "3000"
	}

	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"content-type", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           5 * time.Minute}))

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
	})

	m.Get("/games", controllers.GetGames)
	m.Get("/games/(?P<id>[a-zA-Z0-9]{24})", controllers.GetGameById)

	m.Post("/games", controllers.CreateGame)

	m.Get("/users", controllers.GetUsers)

	m.RunOnAddr(host + ":" + port)
}
