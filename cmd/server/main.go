package main

import (
	"fmt"
	"net/http"

	"github.com/manoj-negi/production-ready-api/internal/database"
	transportHttp "github.com/manoj-negi/production-ready-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {

	var err error

	_, err = database.NewDatabase()

	if err != nil {
		return err
	}

	handler := transportHttp.NewHandler()
	handler.SetupRouter()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("here we go")
	}
}
