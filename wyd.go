package main

import "github.com/alecthomas/kong"

type Wyd struct {
	Version kong.VersionFlag `short:"v" help:"Print the version number"`
}
