package client

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/log"
)

var (
	q   = db.Query
	ctx = db.Ctx
)

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run() error {
	l := log.Get()

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

	_, err := q.CreateClient(ctx, name)
	if err != nil {
		l.Fatal(err)
	}

	return nil
}
