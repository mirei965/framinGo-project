package main

import (
	"log"
	"myapp/data"
	"myapp/handlers"
	"myapp/middleware"
	"os"

	"github.com/mirei965/framinGo"
)

func InitApplication() *application {
	path, err := os.Getwd() //get workind directory
	if err != nil {
		log.Fatal(err)
	}

	//init framingo
	fra := &framinGo.FraminGo{}
	err = fra.New(path)
	if err != nil {
		log.Fatal(err)
	}

	fra.AppName = "myapp"

	myMiddleware := &middleware.Middleware{
		App: fra,
	}

	myHandlers := &handlers.Handlers{
		App: fra,
	}

	app := &application{
		App:        fra,
		Handlers:   myHandlers,
		Middleware: myMiddleware,
	}
	app.App.Routes = app.routes()

	app.Models = data.New(app.App.DB.Pool)
	myHandlers.Models = *app.Models
	app.Middleware.Models = *app.Models

	return app

}
