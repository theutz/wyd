package cli

import (
	"github.com/theutz/wyd/internal/cli/cmds/clients"
	"github.com/theutz/wyd/internal/cli/cmds/projects"
)

type RootCmd struct {
	Clients      clients.ClientsCmd   `cmd:"" help:"working with clients" aliases:"client,c"`
	Projects     projects.ProjectsCmd `cmd:"" help:"working with projects" aliases:"project,p"`
	Debug        bool                 `short:"v" name:"verbose" help:"enable verbose logging"`
	DatabasePath string               `short:"d" help:"where to store the database" type:"existingfile"`
}
