package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

func initLogger() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetPrefix("wyd")
	return logger
}

func initArgs() []string {
	args := os.Args[1:]
	return args
}

func initApp() app.Application {
	params := app.NewAppParams{
		Logger: initLogger(),
		Args:   initArgs(),
	}

	app := app.NewApp(params)
	return app
}

func main() {
	app := initApp()

	err := app.Run()
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		app.Logger().Error(err)
		app.Exit(1)
	}
}
