package bindings

import (
	"context"
	"database/sql"
	"embed"
	"os"

	"github.com/adrg/xdg"
	"github.com/charmbracelet/log"
	"github.com/pressly/goose/v3"
	"github.com/theutz/wyd/internal/queries"
)

type Bindings struct {
	DebugLevel DebugLevel
	Db         *sql.DB
	DbFile     string
	Context    context.Context
	Logger     *log.Logger
	Queries    queries.Queries
}

type DebugLevel int

func (d DebugLevel) AfterApply(b Bindings) error {
	var l log.Level

	switch d {
	case 1:
		l = log.InfoLevel
	case 2:
		l = log.DebugLevel
	default:
		l = log.WarnLevel
	}

	b.Logger.SetLevel(l)

	return nil
}

var ddl string

func (b *Bindings) initLogger() {
	b.Logger = log.New(os.Stderr)
	b.Logger.SetPrefix("wyd")
	b.Logger.SetLevel(log.WarnLevel)
}

func (b *Bindings) initDb() {
	b.Context = context.Background()
	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		b.Logger.Fatal(err)
	}
	b.DbFile = db_file

	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(b.Context, ddl); err != nil {
		log.Fatal(err)
	}
	b.Db = db
}

func (b *Bindings) initMigrations(fs embed.FS) {
	goose.SetBaseFS(fs)

	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}

	if err := goose.Up(b.Db, "migrations"); err != nil {
		log.Fatal(err)
	}
}

func (b *Bindings) initQueries() {
	b.Queries = *queries.New(b.Db)
}

func Init(fs embed.FS) Bindings {
	b := Bindings{}
	b.initLogger()
	b.initDb()
	b.initMigrations(fs)
	b.initQueries()

	return b
}
