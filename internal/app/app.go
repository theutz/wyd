package app

import (
	"embed"
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/config"
)

var logger = log.New(os.Stderr)

func init() {
	logger.SetPrefix("app")
}

type Application interface {
	Exit(code int)
	ExitCode() int
	Args() []string
	Run() error
}

type App struct {
	args        []string
	exitCode    int
	config      config.Config
	migrationFS embed.FS
	isFatal     bool
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
	config, err := config.NewConfig()
	if err != nil {
		logger.Warnf("loading config")
		return err
	}

	parser, err := kong.New(
		&cli,
		kong.Name("wyd"),
		kong.Description("whatcha doing? a time tracking helper"),
		kong.Exit(a.Exit),
		kong.UsageOnError(),
		kong.Bind(config),
		kong.BindTo(a, (*Application)(nil)),
	)
	if err != nil {
		logger.Warn("creating parser", "parser", parser)
		return err
	}

	context, err := parser.Parse(a.Args())
	if err != nil {
		logger.Warn("parsing args", "args", a.Args())
		return err
	}

	err = context.Run()
	context.FatalIfErrorf(err)

	return err
}

func (a *App) Config() config.Config {
	return a.config
}

type NewAppParams struct {
	Args           []string
	Config         *config.Config
	IsFatalOnError *bool
	MigrationsFS   *embed.FS
}

func NewApp(params NewAppParams) Application {
	if params.IsFatalOnError == nil {
		b := bool(true)
		params.IsFatalOnError = &b
	}

	if params.Args == nil {
		params.Args = os.Args[1:]
	}

	if params.Config == nil {
		var err error
		params.Config, err = config.NewConfig()
		if err != nil {
			logger.Error(err)
		}
	}

	if params.MigrationsFS == nil {
		logger.Fatal("embedded migrations not loaded")
	}

	app := &App{
		args:        params.Args,
		config:      *params.Config,
		isFatal:     *params.IsFatalOnError,
		migrationFS: *params.MigrationsFS,
		exitCode:    8,
	}

	return app
}
