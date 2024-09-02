package main

import (
	"context"
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/theutz/wyd/internal/exit"
	"github.com/theutz/wyd/queries"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc@latest generate

const shaLen = 7

var (
	// Version contains the appliction version number. It's set via ldflags when
	// building.
	Version = ""

	// CommitSHA contains the SHA of the commit that this application was built
	// against. It's set via ldflags when building.
	CommitSHA = ""
)

func exiter(code int) {
	os.Exit(code)
}

type Context struct {
	log     *log.Logger
	queries *queries.Queries
	dbCtx   context.Context
}

func getVersion() string {
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

	return version
}

func makeLogger(w io.Writer) *log.Logger {
	l := log.New(w)
	log.SetLevel(log.WarnLevel)
	log.SetPrefix("wyd")
	return l
}

func run(stdout io.Writer, stderr io.Writer, exiter func(int)) error {
	log := makeLogger(stdout)
	db, err := initDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = setupGoose(log, db)
	if err != nil {
		log.Fatal(err)
	}

	context := &Context{
		log:     log,
		queries: queries.New(db),
		dbCtx:   dbContext,
	}

	cli := &Wyd{}
	ctx := kong.Parse(
		cli,
		kong.Description("Whatch'ya doin'?"),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact:             false,
			Summary:             false,
			NoExpandSubcommands: true,
		}),
		kong.Writers(stdout, stderr),
		kong.Exit(exiter),
		kong.Vars{
			"version": getVersion(),
			"db_file": dbFile,
		})
	return ctx.Run(context)
}

func main() {
	err := run(os.Stdin, os.Stderr, exiter)
	if err != nil {
		if errors.Is(err, exit.ErrAborted) ||
			errors.Is(err, huh.ErrUserAborted) {
			log.Warn(err)
			os.Exit(exit.StatusAborted)
		}
		log.Error(err)
		os.Exit(1)
	}
}
