package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("main")
}

func main() {
	params := app.NewAppParams{}
	app := app.NewApp(params)

	err := app.Run()
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		logger.Error(err)
		app.Exit(1)
	}

	app.Exit(0)
}
