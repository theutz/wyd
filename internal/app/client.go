package app

import (
	"fmt"

	"github.com/theutz/wyd/internal/queries/clients"
)

type ClientCmd struct {
	List ClientListCmd `cmd:"" help:"list all clients"`
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
