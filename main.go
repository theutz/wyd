package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

var logger = log.New(os.Stderr)

//go:embed internal/migrations/*.sql
var embeddedMigrations embed.FS

func init() {
	logger.SetPrefix("main")
}

func main() {
	params := app.NewAppParams{
		MigrationsFS: &embeddedMigrations,
	}
	app := app.NewApp(params)

	err := app.Run()
	if err != nil {
		err = fmt.Errorf("error: %w", err)
		logger.Error(err)
		app.Exit(1)
	}

	app.Exit(0)
}
