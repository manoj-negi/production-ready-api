package main

import "fmt"

type App struct {
}

func (app *App) Run() error {
	fmt.Println("inside Run")
	return nil
}

func main() {
	fmt.Println("----start-")

	app := App{}

	if err := app.Run(); err != nil {
		fmt.Println("here we go")
	}
}
