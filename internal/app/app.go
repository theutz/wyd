package app

import (
	"fmt"

	"github.com/charmbracelet/log"
)

type Application interface {
	// Exit()
	// ExitCode() int
	Logger() *log.Logger
	Args() []string
}

type App struct {
	logger *log.Logger
	args   []string
}

func (a *App) Logger() *log.Logger {
	return a.logger
}

func (a *App) Args() []string {
	return a.args
}

type NewAppParams struct {
	Logger *log.Logger
	Args   []string
}

func NewApp(params NewAppParams) Application {
	if params.Logger == nil {
		fmt.Printf("wyd: no logger provided")
	}

	if params.Args == nil {
		params.Logger.Fatalf("no args provided")
	}

	a := &App{
		logger: params.Logger,
		args:   params.Args,
	}

	return a
}
