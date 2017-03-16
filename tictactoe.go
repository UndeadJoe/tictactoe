package main

import (
	"github.com/go-martini/martini"
	"./controllers"
	"net/http"
)

func main() {
	m := martini.Classic()

	m.Use(func(w http.ResponseWriter) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Credentials", "true")
	})

	m.Get("/games", controllers.GetGames)
	m.Get("/games/:id", controllers.GetGameById)

	m.Get("/users", controllers.GetUsers)

	m.Run()
}