package projects

import "fmt"

type ListCmd struct{}

type AddCmd struct{}

type ProjectsCmd struct {
	Add  AddCmd  `cmd:"" help:"add a project"`
	List ListCmd `cmd:"" default:"withargs" help:"list all projects"`
}

func (cmd *ListCmd) Run() error {
	fmt.Println("list all projects")
	return nil
}
