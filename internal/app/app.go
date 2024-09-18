package app

import (
	"context"
	"embed"
	"fmt"
	"os"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/config"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/queries/clients"
	"github.com/theutz/wyd/internal/queries/projects"
)

type Application interface {
	Exit(code int)
	ExitCode() int
	Args() []string
	Run() error
}

type App struct {
	config      config.Config
	migrationFS embed.FS
	args        []string
	isFatal     bool
	exitCode    int
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
	ctx := context.Background()
	config := a.Config()

	connection, err := db.NewConnection(
		ctx,
		a.migrationFS,
		a.Config().DatabasePath,
	)
	if err != nil {
		return fmt.Errorf("creating db: %w", err)
	}
	defer connection.Close()

	clients := clients.New(connection)
	projects := projects.New(connection)

	var cli Cli

	parser, err := kong.New(
		&cli,
		kong.Name("wyd"),
		kong.Description("whatcha doing? a time tracking helper"),
		kong.Exit(a.Exit),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             true,
			FlagsLast:           true,
			NoExpandSubcommands: true,
			Tree:                true,
			NoAppSummary:        false,
			Summary:             true,
			Indenter:            nil,
			WrapUpperBound:      -1,
		}),
		kong.Bind(
			config,
			a,
			clients,
			projects,
		),
		kong.BindTo(ctx, (*context.Context)(nil)),
	)
	if err != nil {
		return fmt.Errorf("creating parser: %w", err)
	}

	kctx, err := parser.Parse(a.Args())
	if err != nil {
		return fmt.Errorf("parsing args: %w", err)
	}

	err = kctx.Run()
	if err != nil {
		return fmt.Errorf("while running kong: %w", err)
	}

	kctx.FatalIfErrorf(err)

	return nil
}

func (a *App) Config() config.Config {
	return a.config
}

type NewAppParams struct {
	Config         *config.Config
	IsFatalOnError *bool
	MigrationsFS   *embed.FS
	Context        *context.Context
	Args           []string
}

func NewApp(params NewAppParams) Application {
	logger := log.WithPrefix("app")

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
		exitCode:    8,
	}

	return app
}
