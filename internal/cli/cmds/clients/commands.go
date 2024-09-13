package clients

import (
	"errors"
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

func (cmd *ListCmd) Run(app *app.Context) error {
	db := app.Db()
	q := queries.New(db)

	clients, err := q.ListClients(app.Ctx())
	if err != nil {
		return err
	}

	cmd.clients = clients

	if err := cmd.Print(); err != nil {
		return err
	}

	return nil
}

func (cmd *ListCmd) Print() error {
	if len(cmd.clients) < 1 {
		return errors.New("no clients found")
	}

	headers := []string{"ID", "Name"}
	rows := [][]string{}

	for _, c := range cmd.clients {
		id := strconv.Itoa(int(c.ID))
		row := []string{id, c.Name}
		rows = append(rows, row)
	}

	t := out.Table(headers, rows)
	fmt.Println(t)

	return nil
}

func (cmd *DeleteCmd) Run(app *app.Context) error {
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
