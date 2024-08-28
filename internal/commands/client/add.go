package client

import (
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
)

var (
	q   = db.Query
	ctx = db.Ctx
)

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run() error {
	log.Debug("adding client")

	name := cmd.Name
	log.Debug("flag", "name", name)

	if name == "" {
		err := huh.NewInput().
			Title("Name").
			Inline(true).
			Value(&name).
			Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Debug("input", "name", name)

	_, err := q.CreateClient(ctx, name)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
