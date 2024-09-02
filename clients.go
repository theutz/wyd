package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
)

type ClientCmd struct {
	Add    AddClientCmd    `cmd:"" help:"add a client"`
	List   ListClientsCmd  `cmd:"" help:"list clients"`
	Remove RemoveClientCmd `cmd:"" help:"remove a client"`
}

func (cmd *ClientCmd) Run() error {
	return nil
}

type RemoveClientCmd struct{}

func (cmd *RemoveClientCmd) Run() error {
	return nil
}

type AddClientCmd struct {
	Name string `short:"n" help:"client name"`
}

func (cmd *AddClientCmd) Run(c *Context) error {
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

	_, err := c.queries.CreateClient(c.dbCtx, name)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

type ListClientsCmd struct{}

func (cmd *ListClientsCmd) Run(c *Context) error {
	log.Debug("listing clients")
	clients, err := c.queries.ListClients(c.dbCtx)
	if err != nil {
		log.Fatal(err)
	}
	t := table.New().
		Headers("Client Name")
	for _, client := range clients {
		t.Row(client.Name)
	}
	fmt.Println(t)
	return nil
}
