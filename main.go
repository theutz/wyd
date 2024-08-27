package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	_ "github.com/mattn/go-sqlite3"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/exit"
	"github.com/theutz/wyd/internal/log"
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

var l = log.Get()

func main() {
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
		kong.Vars{
			"version": version,
			"db_file": db.DbFile,
		})

	if err := ctx.Run(wyd); err != nil {
		if errors.Is(err, exit.ErrAborted) ||
			errors.Is(err, huh.ErrUserAborted) {
			l.Warn(err)
			os.Exit(exit.StatusAborted)
		}
		l.Error(err)
		os.Exit(1)
	}
}
