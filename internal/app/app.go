package app

import (
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/config"
)

type Application interface {
	Exit(code int)
	ExitCode() int
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
	isFatal  bool
}

func (a *App) Logger() *log.Logger {
	return a.logger.WithPrefix("app")
}

func (a *App) Args() []string {
	return a.args
}

func (a *App) Exit(code int) {
	a.exitCode = code
	if a.isFatal {
		os.Exit(code)
	}
}

func (a *App) ExitCode() int {
	return a.exitCode
}

func (a *App) Run() error {
	parser, err := kong.New(
		&cli,
		kong.Name("wyd"),
		kong.Description("whatcha doing? a time tracking helper"),
		kong.Exit(a.Exit),
		kong.UsageOnError(),
	)
	if err != nil {
		a.logger.Warn("creating parser", "parser", parser)
		return err
	}

	context, err := parser.Parse(a.Args())
	if err != nil {
		a.logger.Warn("parsing args", "args", a.Args())
		return err
	}

	err = context.Run(a)

	return err
}

func (a *App) Config() config.Config {
	return a.config
}

type NewAppParams struct {
	Logger         *log.Logger
	Args           []string
	Config         config.Config
	IsFatalOnError bool
}

func NewApp(params NewAppParams) Application {
	if params.Logger == nil {
		fmt.Printf("wyd: no logger provided")
	}

	if params.Args == nil {
		params.Logger.Fatalf("no args provided")
	}

	var fatalOnError bool
	switch params.IsFatalOnError {
	case true:
		fatalOnError = true
	case false:
		fatalOnError = false
	default:
		fatalOnError = true
	}

	app := &App{
		logger:  params.Logger,
		args:    params.Args,
		config:  params.Config,
		isFatal: fatalOnError,
	}

	return app
}
