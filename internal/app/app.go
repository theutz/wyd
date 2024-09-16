package app

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/config"
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
	Config() config.Config
}

type App struct {
	logger   *log.Logger
	args     []string
	exitCode int
	config   config.Config
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

func (a *App) Config() config.Config {
	return a.config
}

type NewAppParams struct {
	Logger *log.Logger
	Args   []string
	Config config.Config
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
		config: params.Config,
	}

	return app
}
