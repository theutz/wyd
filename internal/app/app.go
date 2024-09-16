package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

type Exiter interface {
	Exit(code int)
	ExitCode() int
}

type Application interface {
	Exiter
	Logger() *log.Logger
	Args() []string
	Run() error
}

type App struct {
	logger   *log.Logger
	args     []string
	exitCode int
}

func (a *App) Logger() *log.Logger {
	return a.logger
}

func (a *App) Args() []string {
	return a.args
}

func (a *App) Exit(code int) {
	a.exitCode = code
	os.Exit(code)
}

func (a *App) ExitCode() int {
	return a.exitCode
}

func (a *App) Run() error {
	return nil
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

	app := &App{
		logger: params.Logger,
		args:   params.Args,
	}

	return app
}
