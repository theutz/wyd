package main

import (
	"github.com/alecthomas/kong"
	"github.com/theutz/wyd/internal/bindings"
	"github.com/theutz/wyd/internal/client"
	"github.com/theutz/wyd/internal/project"
)

type Globals struct {
	DbPath     string              `short:"f" help:"set the path for the sqlite database" type:"path" default:"${db_file}" placeholder:"test.db"`
	DebugLevel bindings.DebugLevel `short:"d" help:"Set the debug level" enum:"0,1,2" default:"0"`
}

type Wyd struct {
	Version kong.VersionFlag   `short:"v" help:"Print the version number"`
	Client  client.ClientCmd   `cmd:"" help:"work with clients" aliases:"clients,c"`
	Project project.ProjectCmd `cmd:"" help:"work with projects" aliases:"projects,p"`
	Globals
}
