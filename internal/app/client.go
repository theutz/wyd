package app

import (
	"context"
	"fmt"

	q "github.com/theutz/wyd/internal/queries/clients"
	"github.com/theutz/wyd/internal/views"
)

type ClientCmd struct {
	List   ClientListCmd   `aliases:"show,ls"   cmd:"" help:"list all clients"`
	Add    ClientAddCmd    `aliases:"create,a"  cmd:"" help:"add a new client"`
	Remove ClientRemoveCmd `aliases:"delete,rm" cmd:"" help:"remove a client"`
}

type ClientListCmd struct{}

func (cmd *ClientListCmd) Run(ctx context.Context, q *q.Queries) error {
	clients, err := q.All(ctx)
	if err != nil {
		return fmt.Errorf("loading all clients: %w", err)
	}

	fmt.Println(clients)

	return nil
}

type ClientAddCmd struct {
	Name string `help:"name of the client" required:"" short:"n"`
}

func (cmd *ClientAddCmd) Run(ctx context.Context, q *q.Queries) error {
	client, err := q.Create(ctx, cmd.Name)
	if err != nil {
		return fmt.Errorf("creating client: %w", err)
	}

	view := views.Record(client.ToMap())
	fmt.Println(view)

	return nil
}

type ClientRemoveCmd struct {
	Name string `help:"the name of the client" required:"" short:"n"`
}

func (cmd *ClientRemoveCmd) Run(ctx context.Context, q *q.Queries) error {
	client, err := q.DeleteByName(ctx, cmd.Name)
	if err != nil {
		return fmt.Errorf("deleting client by name %s: %w", cmd.Name, err)
	}

	fmt.Println(views.Record(client.ToMap()))

	return nil
}
