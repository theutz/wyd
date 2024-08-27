package db

import (
	"context"
	"database/sql"
	"embed"

	"github.com/adrg/xdg"
	"github.com/pressly/goose/v3"
	"github.com/theutz/wyd/internal/db/queries"
	"github.com/theutz/wyd/internal/log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

var Ctx = context.Background()

var DbFile string

var ddl string

var Db *sql.DB

var Query queries.Queries

func Init() *sql.DB {
	l := log.Get()
	var err error
	DbFile, err = xdg.DataFile("wyd/wyd.db")
	if err != nil {
		l.Fatal(err)
	}

	Db, err := sql.Open("sqlite3", DbFile)
	if err != nil {
		l.Fatal(err)
	}

	if _, err := Db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		l.Fatal(err)
	}

	if _, err := Db.ExecContext(Ctx, ddl); err != nil {
		l.Fatal(err)
	}

	goose.SetBaseFS(embedMigrations)

	switch l.GetLevel().String() {
	case "info":
		goose.SetLogger(l)
		fallthrough
	case "debug":
		goose.SetVerbose(true)
		goose.SetLogger(goose.NopLogger())
	default:
		goose.SetLogger(goose.NopLogger())
	}

	if err := goose.SetDialect("sqlite"); err != nil {
		l.Fatal(err)
	}

	if err := goose.Up(Db, "migrations"); err != nil {
		l.Fatal(err)
	}

	Query = *queries.New(Db)

	return Db
}
