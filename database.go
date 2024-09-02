package main

import (
	"context"
	"database/sql"
	"embed"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/log"
	"github.com/pressly/goose/v3"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var ddl string

var dbContext = context.Background()

var dbFile string

func initDatabase() (*sql.DB, error) {
	dbFile, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		return nil, err
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		return nil, err
	}

	if _, err := db.ExecContext(dbContext, ddl); err != nil {
		return nil, err
	}

	return db, nil
}

func setupGoose(log *log.Logger, db *sql.DB) error {
	goose.SetBaseFS(embedMigrations)

	switch log.GetLevel().String() {
	case "info":
		goose.SetLogger(log.StandardLog())
		fallthrough
	case "debug":
		goose.SetVerbose(true)
		goose.SetLogger(goose.NopLogger())
	default:
		goose.SetLogger(goose.NopLogger())
	}

	if err := goose.SetDialect("sqlite"); err != nil {
		return err
	}

	if err := goose.Up(db, "migrations"); err != nil {
		return err
	}

	return nil
}
