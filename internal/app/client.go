package app

import (
	"fmt"

	"github.com/theutz/wyd/internal/queries/clients"
)

type ClientCmd struct {
	List ClientListCmd `cmd:"" aliases:"show,ls" help:"list all clients"`
	Add  ClientAddCmd  `cmd:"" aliases:"create,a" help:"add a new client"`
}

type ClientListCmd struct{}

func (cmd *ClientListCmd) Run(app *App, c *clients.Queries) error {
	clients, err := c.All(app.Context())
	if err != nil {
		return err
	}
	fmt.Println(clients)
	return nil
}

type ClientAddCmd struct {
	Name string `short:"n" help:"name of the client"`
}

func (cmd *ClientAddCmd) Run(app *App, c *clients.Queries) error {
	client, err := c.Create(app.Context(), cmd.Name)
	if err != nil {
		return err
	}
	fmt.Println(client)
	return nil
}
