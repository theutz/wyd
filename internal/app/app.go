package app

import (
	"context"
	"embed"
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/config"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/queries/clients"
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
	Context() context.Context
}

type App struct {
	args        []string
	exitCode    int
	config      config.Config
	migrationFS embed.FS
	isFatal     bool
	context     context.Context
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

func (a *App) Context() context.Context {
	return a.context
}

func (a *App) Run() error {
	config := a.Config()

	db, err := db.NewDb(
		a.Context(),
		a.migrationFS,
		a.Config().DatabasePath,
	)
	if err != nil {
		logger.Warn("creating db")
		return err
	}

	clients := clients.New(db)

	parser, err := kong.New(
		&cli,
		kong.Name("wyd"),
		kong.Description("whatcha doing? a time tracking helper"),
		kong.Exit(a.Exit),
		kong.UsageOnError(),
		kong.Bind(
			config,
			a,
			clients,
		),
	)
	if err != nil {
		logger.Warn("creating parser", "parser", parser)
		return err
	}

	kctx, err := parser.Parse(a.Args())
	if err != nil {
		logger.Warn("parsing args", "args", a.Args())
		return err
	}

	err = kctx.Run()
	kctx.FatalIfErrorf(err)

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
	Context        *context.Context
}

func NewApp(params NewAppParams) Application {
	if params.IsFatalOnError == nil {
		b := bool(true)
		params.IsFatalOnError = &b
	}

	if params.Args == nil {
		params.Args = os.Args[1:]
	}

	if params.Context == nil {
		c := context.Background()
		params.Context = &c
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
		context:     *params.Context,
		exitCode:    8,
	}

	return app
}
