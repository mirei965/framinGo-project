package main

import (
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"

	"github.com/mirei965/framinGo"
)

type application struct {
	App        *framinGo.FraminGo
	Handlers   *handlers.Handlers
	Models     *data.Models
	Middleware *middleware.Middleware
}

func main() {
	f := InitApplication()
	f.App.ListenAndServe()

}
