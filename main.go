package main

import (
	_ "embed"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
	"github.com/theutz/wyd/internal/config"
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

//go:embed default-config.yml
var defaultConfig []byte

func initConfig() (config.Config, error) {
	c, err := config.NewConfig(defaultConfig)
	if err != nil {
		return nil, fmt.Errorf("while initializing config: %w", err)
	}

	return c, nil
}

func initApp() app.Application {
	l := initLogger()
	params := app.NewAppParams{
		Logger: l,
		Args:   initArgs(),
	}

	config, err := initConfig()
	if err != nil {
		l.Fatal(err)
	}
	params.Config = config

	app := app.NewApp(params)
	return app
}

func main() {
	// TODO: make --print-config flag
	app := initApp()

	err := app.Run()
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		app.Logger().Error(err)
		app.Exit(1)
	}

	app.Exit(0)
}
