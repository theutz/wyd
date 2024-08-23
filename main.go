package main

import (
	"errors"
	"fmt"
	"os"
	"runtime/debug"

	"github.com/adrg/xdg"
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/internal/exit"
	"github.com/theutz/wyd/internal/logger"
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

func main() {
	log := logger.New(false)

	db_file, err := xdg.DataFile("wyd/wyd.db")
	if err != nil {
		log.Fatal(err)
	}

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
		})
	if err := ctx.Run(); err != nil {
		if errors.Is(err, exit.ErrAborted) || errors.Is(err, huh.ErrUserAborted) {
			os.Exit(exit.StatusAborted)
		}
		fmt.Println(err)
		os.Exit(1)
	}
}
