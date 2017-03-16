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
	})

	m.Get("/games", controllers.GetGames)

	m.Get("/users", controllers.GetUsers)

	m.Run()
}