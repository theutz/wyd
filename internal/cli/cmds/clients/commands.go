package clients

import (
	"errors"
	"fmt"

	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/db/queries"
)

func (cmd *AddClientCmd) Run(app *app.Context) error {
	ctx, q := app.Queries()

	client, err := q.AddClient(ctx, cmd.Name)
	if err != nil {
		return err
	}

	fmt.Println(client.Render())

	return nil
}

func (cmd *ListClientsCmd) Run(app *app.Context) error {
	ctx, q := app.Queries()

	clients, err := q.ListClients(ctx)
	if err != nil {
		return err
	}

	c := queries.Clients(clients)
	fmt.Println(c.Render())

	return nil
}

func (cmd *DeleteClientsCmd) Run(app *app.Context) error {
	ctx, q := app.Queries()

	var client queries.Client
	var err error

	switch {
	case cmd.Id != 0:
		client, err = q.DeleteClient(ctx, cmd.Id)
		if err != nil {
			return fmt.Errorf("error: while deleting client by id: %w", err)
		}
	case cmd.Name != "":
		client, err = q.DeleteClientByName(ctx, cmd.Name)
		if err != nil {
			return fmt.Errorf("error: while deleting client by name: %w", err)
		}
	default:
		return errors.New("error: client not found")
	}

	fmt.Println(client.Render())

	return nil
}
