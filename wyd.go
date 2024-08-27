package main

import (
	"github.com/alecthomas/kong"
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/client"
	"github.com/theutz/wyd/project"
	"github.com/theutz/wyd/tasks"
)

type Globals struct {
	DbPath     string              `short:"f" help:"set the path for the sqlite database" type:"path" default:"${db_file}" placeholder:"test.db"`
	DebugLevel bindings.DebugLevel `short:"d" help:"Set the debug level" enum:"0,1,2" default:"0"`
}

type Wyd struct {
	Version kong.VersionFlag   `short:"v" help:"Print the version number"`
	Client  client.ClientCmd   `cmd:"" help:"work with clients" aliases:"clients,c"`
	Project project.ProjectCmd `cmd:"" help:"work with projects" aliases:"projects,p"`
	Task    tasks.TaskCmd      `cmd:"" help:"work with tasks" aliases:"tasks,t"`
	Globals
}
