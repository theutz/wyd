package app

import (
	"fmt"

	"github.com/theutz/wyd/internal/queries/clients"
)

type ClientCmd struct {
	List ClientListCmd `aliases:"show,ls"  cmd:"" help:"list all clients"`
	Add  ClientAddCmd  `aliases:"create,a" cmd:"" help:"add a new client"`
}

type ClientListCmd struct{}

func (cmd *ClientListCmd) Run(app *App, c *clients.Queries) error {
	clients, err := c.All(app.Context())
	if err != nil {
		err = fmt.Errorf("loading all clients: %w", err)

		return err
	}

	fmt.Println(clients)

	return nil
}

type ClientAddCmd struct {
	Name string `help:"name of the client" short:"n"`
}

func (cmd *ClientAddCmd) Run(app *App, c *clients.Queries) error {
	client, err := c.Create(app.Context(), cmd.Name)
	if err != nil {
		err = fmt.Errorf("creating client: %w", err)

		return err
	}

	fmt.Println(client)

	return nil
}
