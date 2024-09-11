package clients

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/theutz/wyd/internal/cli/app"
	"github.com/theutz/wyd/internal/cli/out"
	"github.com/theutz/wyd/internal/db/clients"
)

func (cmd *ListCmd) Run(app *app.Context) error {
	db := app.GetDb()
	q := clients.New(db)

	clients, err := q.ListClients(app.GetCtx())
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
