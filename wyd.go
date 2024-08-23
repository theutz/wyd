package main

import (
	"github.com/alecthomas/kong"
	"github.com/theutz/wyd/client"
	"github.com/theutz/wyd/project"
)

type Globals struct {
	DbPath  string           `short:"f" help:"set the path for the sqlite database" type:"path" default:"${db_file}" placeholder:"test.db"`
	Version kong.VersionFlag `short:"v" help:"Print the version number"`
	Debug   bool             `short:"d" help:"Turn on verbose logging"`
}

type Wyd struct {
	Client  client.ClientCmd   `cmd:"" help:"work with clients" alias:"clients,c"`
	Project project.ProjectCmd `cmd:"" help:"work with projects" alias:"projects,p"`
	Globals
}
