package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type App struct {
	logger *log.Logger
}

func NewApp(app *App) {
	fmt.Println("hello there")
}

func main() {
	logger := log.New(os.Stderr)
	app := &App{
		logger: logger,
	}

	NewApp(app)
}
