package client

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/internal/bindings"
)

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run(b bindings.Bindings) error {
	b.Logger.Debug("adding client")
	name := cmd.Name
	b.Logger.Debug("flag", "name", name)

	if name == "" {
		err := huh.NewInput().
			Title("Name").
			Inline(true).
			Value(&name).
			Run()
		if err != nil {
			b.Logger.Fatal(err)
		}
	}
	b.Logger.Debug("input", "name", name)

	_, err := b.Queries.CreateClient(b.Context, name)
	if err != nil {
		b.Logger.Fatal(err)
	}

	return nil
}
