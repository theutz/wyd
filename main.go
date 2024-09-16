package main

import (
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

func main() {
	logger := log.New(os.Stderr)

	params := app.NewAppParams{
		Logger: logger,
	}

	app.NewApp(params)
}
