package app

import "github.com/charmbracelet/log"

type Application interface {
	// Exit()
	// ExitCode() int
	Logger() *log.Logger
	// Args() []string
}

type App struct {
	logger *log.Logger
}

func (a *App) Logger() *log.Logger {
	return a.logger
}

type NewAppParams struct {
	Logger *log.Logger
}

func NewApp(params NewAppParams) Application {
	a := &App{
		logger: params.Logger,
	}

	return a
}
