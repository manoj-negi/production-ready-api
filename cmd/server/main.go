package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/manoj-negi/production-ready-api/internal/comment"
	"github.com/manoj-negi/production-ready-api/internal/database"
	transportHttp "github.com/manoj-negi/production-ready-api/internal/transport/http"
)

type App struct {
}

func (app *App) Run() error {

	var err error

	store, err := database.NewDatabase()

	if err != nil {
		fmt.Println("Failed to connect", err)
		return err
	}

	cmtService := comment.NewService(store)

	cmtService.PostComment(context.Background(), comment.Comment{
		ID:     "0f87f266-a337-440f-a2fd-5a92acf6d8c0",
		Slug:   "manual-test",
		Author: "Sonam dubey",
		Body:   "Hello World",
	})

	fmt.Println(cmtService.GetComment(context.Background(), "01fd02a5-1ace-48b5-8822-ee3795a7d6e9"))

	handler := transportHttp.NewHandler()
	handler.SetupRouter()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		fmt.Println("Failed to setup server")
		return err
	}
	fmt.Println("succesfully connected databse", store)
	return nil
}

func main() {

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("here we go")
	}
}
