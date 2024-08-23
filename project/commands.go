package project

import "github.com/charmbracelet/log"

type ProjectCmd struct {
	Add    AddCmd    `cmd:"" help:"add a project"`
	List   ListCmd   `cmd:"" help:"list all projects"`
	Remove RemoveCmd `cmd:"" help:"remove a project"`
}

func (cmd *ProjectCmd) Run(log *log.Logger) error {
	return nil
}

type AddCmd struct{}

func (cmd *AddCmd) Run() error {
	return nil
}

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	return nil
}

type RemoveCmd struct{}

func (cmd *RemoveCmd) Run() error {
	return nil
}
