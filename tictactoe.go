package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/cors"
	"./controllers"
	"time"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(cors.Allow(&cors.Options{
		AllowOrigins:     []string{"http://localhost"},
		AllowMethods:     []string{"OPTIONS", "GET", "POST"},
		AllowHeaders:     []string{"content-type", "x-token"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge: 	5 * time.Minute}))

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "tapplication/json; charset=utf-8")
	})

	m.Get("/games", controllers.GetGames)
	m.Get("/games/(?P<id>[a-zA-Z0-9]{24})", controllers.GetGameById)

	m.Post("/games", controllers.CreateGame)

	m.Get("/users", controllers.GetUsers)

	m.Run()
}