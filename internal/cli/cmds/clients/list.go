package clients

import (
	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/data/clients"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(app *app.Context) error {
	db := app.GetDb()
	q := clients.New(db)

	clients, err := q.ListClients(app.GetCtx())
	if err != nil {
		return err
	}

	err = printClients(clients)
	if err != nil {
		return err
	}

	return nil
}
