package client

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/bindings"
)

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run(b bindings.Bindings) error {
	l.Debug("adding client")
	name := cmd.Name
	l.Debug("flag", "name", name)

	if name == "" {
		err := huh.NewInput().
			Title("Name").
			Inline(true).
			Value(&name).
			Run()
		if err != nil {
			l.Fatal(err)
		}
	}
	l.Debug("input", "name", name)

	_, err := b.Queries.CreateClient(b.Context, name)
	if err != nil {
		l.Fatal(err)
	}

	return nil
}
