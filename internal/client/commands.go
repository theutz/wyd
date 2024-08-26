package client

import (
	"context"

	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name string `gorm:"unique,not null"`
}

type ClientCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" help:"list clients"`
	Remove RemoveCmd `cmd:"" help:"remove a client"`
}

func (cmd *ClientCmd) Run() error {
	return nil
}

type AddCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddCmd) Run(log *clog.Logger, q *queries.Queries) error {
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

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	return nil
}

type RemoveCmd struct{}

func (cmd *RemoveCmd) Run() error {
	return nil
}
