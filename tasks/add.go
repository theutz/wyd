package tasks

import (
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/queries"
)

type AddCmd struct {
	Name    string `short:"n" help:"the name of the task"`
	Project string `short:"p" help:"the name of the project"`
	p       queries.CreateTaskParams
	b       bindings.Bindings
}

func (cmd *AddCmd) handleName() {
}

func (cmd *AddCmd) Run(b bindings.Bindings) error {
	cmd.b = b
	cmd.p = queries.CreateTaskParams{}

	cmd.handleName()

	return nil
}
