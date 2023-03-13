package main

import (
	"fmt"
	"net/http"

	transportHttp "github.com/manoj-negi/production-ready-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {
	fmt.Println("setting up our apps")

	handler := transportHttp.NewHandler()
	handler.SetupRouter()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}

	return nil
}

func main() {
	fmt.Println("----start-")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("here we go")
	}
}
