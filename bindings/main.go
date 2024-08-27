package bindings

import (
	"context"
	"database/sql"
	"embed"

	"github.com/adrg/xdg"
	clog "github.com/charmbracelet/log"
	"github.com/pressly/goose/v3"
	"github.com/theutz/wyd/internal/log"
	"github.com/theutz/wyd/queries"
)

type Bindings struct {
	DebugLevel DebugLevel
	Db         *sql.DB
	DbFile     string
	Context    context.Context
	Queries    queries.Queries
}

type DebugLevel int

var l = log.Get()

func (d DebugLevel) AfterApply(b Bindings) error {
	var lvl clog.Level

	switch d {
	case 1:
		lvl = clog.InfoLevel
	case 2:
		lvl = clog.DebugLevel
	default:
		lvl = clog.WarnLevel
	}

	l.SetLevel(lvl)

	return nil
}

var ddl string

func (b *Bindings) initDb() {
	l := log.Get()
	b.Context = context.Background()
	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		l.Fatal(err)
	}
	b.DbFile = db_file

	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		clog.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		clog.Fatal(err)
	}

	if _, err := db.ExecContext(b.Context, ddl); err != nil {
		clog.Fatal(err)
	}
	b.Db = db
}

func (b *Bindings) initMigrations(fs embed.FS) {
	l := log.Get()
	goose.SetBaseFS(fs)
	goose.SetLogger(goose.NopLogger())
	switch b.DebugLevel {
	case 1:
		goose.SetLogger(l)
		fallthrough
	case 2:
		goose.SetVerbose(true)
		goose.SetLogger(goose.NopLogger())
	}

	if err := goose.SetDialect("sqlite"); err != nil {
		clog.Fatal(err)
	}

	if err := goose.Up(b.Db, "migrations"); err != nil {
		clog.Fatal(err)
	}
}

func (b *Bindings) initQueries() {
	b.Queries = *queries.New(b.Db)
}

func Init(fs embed.FS) Bindings {
	b := Bindings{}
	b.initDb()
	b.initMigrations(fs)
	b.initQueries()

	return b
}
