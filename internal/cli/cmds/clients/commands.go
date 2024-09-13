package clients

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/cli/out"
	"github.com/theutz/wyd/internal/db/queries"
)

func (cmd *AddClientCmd) Run(app *app.Context) error {
	ctx, q := app.Queries()

	client, err := q.AddClient(ctx, cmd.Name)
	if err != nil {
		return err
	}
	cmd.client = client

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
	case cmd.Name != "":
		client, err = q.DeleteClientByName(ctx, cmd.Name)
	default:
		return errors.New("client not found")
	}

	if err != nil {
		return err
	}
	cmd.client = client

	// TODO: Create renderer
	cmd.Ouptut()

	return nil
}

func (cmd *DeleteClientsCmd) Ouptut() {
	id := strconv.Itoa(int(cmd.client.ID))
	client := map[string]string{
		"ID":   id,
		"Name": cmd.client.Name,
	}
	record := out.Record(client)
	fmt.Println(record)
}
