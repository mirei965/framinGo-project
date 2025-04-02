package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"github.com/mirei965/framingo"
)

type application struct {
	App *framingo.Framingo
	Handlers *handlers.Handlers
	Models *data.Models
	Middleware *middleware.Middleware
}

func main() {
	f := InitApplication()
	f.App.ListenAndServe()

}
