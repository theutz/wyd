package main

import (
	"github.com/alecthomas/kong"
	"github.com/theutz/wyd/client"
	"github.com/theutz/wyd/project"
)

type Wyd struct {
	Project project.ProjectCmd `cmd:"" help:"work with projects"`
	Client  client.ClientCmd   `cmd:"" help:"work with clients"`
	Version kong.VersionFlag   `short:"v" help:"Print the version number"`
}
