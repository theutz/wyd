package main

import (
	"github.com/alecthomas/kong"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/commands/client"
	"github.com/theutz/wyd/internal/commands/entries"
	"github.com/theutz/wyd/internal/commands/project"
	"github.com/theutz/wyd/internal/commands/tasks"
)

type Globals struct {
	DbPath     string           `short:"f" help:"set the path for the sqlite database" type:"path" default:"${db_file}" placeholder:"test.db"`
	DebugLevel DebugLevel       `short:"d" help:"Set the debug level" enum:"0,1,2" default:"0"`
	Version    kong.VersionFlag `short:"v" help:"Print the version number"`
}

type Wyd struct {
	Client  client.ClientCmd   `cmd:"" help:"work with clients" aliases:"clients,c"`
	Project project.ProjectCmd `cmd:"" help:"work with projects" aliases:"projects,p"`
	Task    tasks.TaskCmd      `cmd:"" help:"work with tasks" aliases:"tasks,t"`
	Entry   entries.EntryCmd   `cmd:"" help:"work with entries" aliases:"entries,e"`
	Globals
}

type DebugLevel int

func (d DebugLevel) AfterApply() error {
	switch d {
	case 1:
		log.SetLevel(log.InfoLevel)
	case 2:
		log.SetLevel(log.DebugLevel)
	default:
		log.SetLevel(log.WarnLevel)
	}

	return nil
}
