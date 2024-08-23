package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
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
	logger := log.New(os.Stderr)
	logger.SetPrefix("wyd")
	logger.SetLevel(log.WarnLevel)

	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		logger.Fatal(err)
	}
	db := db.New(db_file, logger)

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
		kong.Bind(db))
	if err := ctx.Run(wyd); err != nil {
		if errors.Is(err, exit.ErrAborted) || errors.Is(err, huh.ErrUserAborted) {
			os.Exit(exit.StatusAborted)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
