package clients

import (
	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/db/clients"
)

func (cmd *AddCmd) Run(app *app.Context) error {
	ctx := app.GetCtx()
	db := app.GetDb()
	defer db.Close()
	q := clients.New(db)

	name := cmd.Name

	client, err := q.AddClient(ctx, name)
	if err != nil {
		return err
	}

	err = printClients([]clients.Client{client})
	if err != nil {
		return err
	}

	return nil
}
