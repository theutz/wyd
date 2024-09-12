package clients

import (
	"fmt"
	"strconv"

	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/cli/out"
	"github.com/theutz/wyd/internal/db/queries"
)

func (cmd *AddCmd) Run(app *app.Context) error {
	ctx := app.Ctx()
	db := app.Db()
	defer db.Close()
	q := queries.New(db)

	name := cmd.Name

	client, err := q.AddClient(ctx, name)
	if err != nil {
		return err
	}
	cmd.client = client

	err = cmd.Output()
	if err != nil {
		return err
	}

	return nil
}

func (cmd *AddCmd) Output() error {
	id := strconv.Itoa(int(cmd.client.ID))
	client := map[string]string{
		"ID":   id,
		"Name": cmd.client.Name,
	}

	record := out.Record(client)
	fmt.Println(record)

	return nil
}
