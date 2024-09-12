package clients

import (
	"fmt"
	"strconv"

	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/cli/out"
)

func (cmd *DeleteCmd) Run(app *app.Context) error {
	ctx, q := app.Queries()

	client, err := q.DeleteClient(ctx, cmd.Name)
	if err != nil {
		return err
	}
	cmd.client = client

	cmd.Ouptut()

	return nil
}

func (cmd *DeleteCmd) Ouptut() {
	id := strconv.Itoa(int(cmd.client.ID))
	client := map[string]string{
		"ID":   id,
		"Name": cmd.client.Name,
	}
	record := out.Record(client)
	fmt.Println(record)
}
