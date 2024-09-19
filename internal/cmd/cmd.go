package cmd

import (
	"github.com/theutz/wyd/internal/cmd/clients"
	"github.com/theutz/wyd/internal/cmd/config"
	"github.com/theutz/wyd/internal/cmd/projects"
)

type Cmd struct {
	Config  config.Cmd          `aliases:""           cmd:"" help:"view wyd configuration"`
	Client  clients.ClientCmd   `aliases:"c,clients"  cmd:"" help:"work with client list"`
	Project projects.ProjectCmd `aliases:"p,projects" cmd:"" help:"work with project list"`
}
