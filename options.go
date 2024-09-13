package main

import (
	"github.com/theutz/wyd/internal/cli/cmds/clients"
	"github.com/theutz/wyd/internal/cli/cmds/projects"
	"github.com/theutz/wyd/internal/cli/cmds/tasks"
)

type RootCmd struct {
	Clients      clients.ClientsCmd   `cmd:"" help:"working with clients" aliases:"client,c"`
	Projects     projects.ProjectsCmd `cmd:"" help:"working with projects" aliases:"project,p"`
	Tasks        tasks.TasksCmd       `cmd:"" help:"working with tasks" aliases:"task,t"`
	Debug        bool                 `short:"v" name:"verbose" help:"enable verbose logging"`
	DatabasePath string               `short:"d" help:"where to store the database" type:"existingfile"`
}
