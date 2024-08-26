package main

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

var (
	db_file string
	ddl     string

	//go:embed db/migrations/*.sql
	embedMigrations embed.FS
)

type Context struct {
	DebugLevel int
	Context    context.Context
	Logger     *log.Logger
	Queries    queries.Queries
}

func (c *Context) initLogger() {
	c.Logger = log.New(os.Stderr)
	c.Logger.SetPrefix("wyd")
	c.Logger.SetLevel(log.WarnLevel)
}

func (c *Context) initDb() *sql.DB {
	c.Context = context.Background()
	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		c.Logger.Fatal(err)
	}

	db, err := sql.Open("sqlite3", db_file)
	if err != nil {
		log.Fatal(err)
	}

	if _, err := db.Exec("PRAGMA foreign_keys = ON"); err != nil {
		log.Fatal(err)
	}

	if _, err := db.ExecContext(c.Context, ddl); err != nil {
		log.Fatal(err)
	}

	return db
}

func (c *Context) initMigrations(db *sql.DB) {
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatal(err)
	}
}

func (c *Context) initQueries(db *sql.DB) {
	c.Queries = *queries.New(db)
}

func initContext() (Context, *sql.DB) {
	c := Context{}
	c.initLogger()
	db := c.initDb()
	c.initMigrations(db)
	c.initQueries(db)

	return c, db
}
