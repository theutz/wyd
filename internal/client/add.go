package client

import (
	"context"

	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run(log *clog.Logger, q *queries.Queries) error {
	log.Debug("adding client")
	name := cmd.Name
	log.Debug("flag", "name", name)

	if name == "" {
		err := huh.NewInput().Title("Name").Inline(true).Value(&name).Run()
		if err != nil {
			log.Fatal(err)
		}
	}
	log.Debug("input", "name", name)

	_, err := q.CreateClient(context.Background(), name)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
