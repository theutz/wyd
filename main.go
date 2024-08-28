package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"runtime/debug"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/exit"
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

func run(stdout io.Writer, stderr io.Writer, exiter func(int)) error {
	log.SetOutput(os.Stderr)
	log.SetLevel(log.WarnLevel)
	log.SetPrefix("wyd")

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

	conn := db.Init()
	defer conn.Close()

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
		kong.Writers(stdout, stderr),
		kong.Exit(exiter),
		kong.Vars{
			"version": version,
			"db_file": db.DbFile,
		})
	return ctx.Run(wyd)
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
