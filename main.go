package main

import (
	"embed"
	"fmt"
	"os"

	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/app"
)

//go:embed internal/migrations/*.sql
var embeddedMigrations embed.FS

func main() {
	logger := log.New(os.Stderr)
	logger.SetPrefix("main")

	params := app.NewAppParams{
		MigrationsFS:   &embeddedMigrations,
		IsFatalOnError: nil,
		Args:           nil,
		Config:         nil,
		Context:        nil,
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
