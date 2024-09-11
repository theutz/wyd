package clients

import (
	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/db/clients"
)

type AddCmd struct {
	Name string `arg:"" help:"the name of the client"`
}

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
