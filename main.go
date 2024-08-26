package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	_ "github.com/mattn/go-sqlite3"
	"github.com/theutz/wyd/internal/exit"
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

	c, db := initContext()
	defer db.Close()

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
		kong.Bind(c))

	if err := ctx.Run(wyd); err != nil {
		if errors.Is(err, exit.ErrAborted) ||
			errors.Is(err, huh.ErrUserAborted) {
			c.Logger.Warn(err)
			os.Exit(exit.StatusAborted)
		}
		c.Logger.Error(err)
		os.Exit(1)
	}
}
