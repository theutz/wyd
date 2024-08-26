package main

import (
	"context"
	"database/sql"
	"embed"
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/pressly/goose/v3"
	"github.com/theutz/wyd/internal/exit"
	"github.com/theutz/wyd/internal/queries"
)

const shaLen = 7

var (
	// Version contains the appliction version number. It's set via ldflags when
	// building.
	Version = ""

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via ldflags when building.
	CommitSHA = ""
)

type debugLevel int

func (d debugLevel) AfterApply(logger *log.Logger) error {
	var l log.Level

	switch d {
	case 1:
		l = log.InfoLevel
	case 2:
		l = log.DebugLevel
	default:
		l = log.WarnLevel
	}

	logger.SetLevel(l)

	return nil
}

var (
	ddl string
	//go:embed db/migrations/*.sql
	embedMigrations embed.FS
)

func main() {
	logger := log.New(os.Stderr)
	logger.SetPrefix("wyd")
	logger.SetLevel(log.WarnLevel)

	qctx := context.Background()
	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		logger.Fatal(err)
	}
	db, err := sql.Open("sqlite3", db_file+";foreign keys=true")
	if _, err := db.ExecContext(qctx, ddl); err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	goose.SetBaseFS(embedMigrations)
	if err := goose.SetDialect("sqlite"); err != nil {
		log.Fatal(err)
	}
	if err := goose.Up(db, "db/migrations"); err != nil {
		log.Fatal(err)
	}
	q := queries.New(db)

	if Version == "" {
		if info, ok := debug.ReadBuildInfo(); ok && info.Main.Sum != "" {
			Version = info.Main.Version
		} else {
			Version = "unknown (built from source)"
		}
	}

	version := fmt.Sprintf("wyd version %s", Version)
	if len(CommitSHA) >= shaLen {
		version += " (" + CommitSHA[:shaLen] + ")"
	}

	wyd := &Wyd{}
	ctx := kong.Parse(
		wyd,
		kong.Description("Whatch'ya doin'?"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             false,
			Summary:             false,
			NoExpandSubcommands: true,
		}),
		kong.Vars{
			"version": version,
			"db_file": db_file,
		},
		kong.Bind(logger),
		kong.Bind(q))
	if err := ctx.Run(wyd); err != nil {
		if errors.Is(err, exit.ErrAborted) || errors.Is(err, huh.ErrUserAborted) {
			os.Exit(exit.StatusAborted)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
