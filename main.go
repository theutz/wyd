package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

func main() {
	logger := log.New(os.Stderr)
	args := os.Args[1:]

	params := app.NewAppParams{
		Logger: logger,
		Args:   args,
	}

	app.NewApp(params)
}
