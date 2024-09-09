package clients

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/internal/cli/context"
	"github.com/theutz/wyd/internal/data/clients"
)

type ListCmd struct{}

type AddCmd struct {
	Name string `arg:"" help:"the name of the client"`
}

func (cmd *AddCmd) Run(app *context.Context) error {
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

type ClientsCmd struct {
	Add  AddCmd  `cmd:"" help:"add a client"`
	List ListCmd `cmd:"" default:"withargs" help:"list all clients"`
}

func printClients(clients []clients.Client) error {
	if len(clients) < 1 {
		return errors.New("no clients found")
	}

	accentColor := lipgloss.ANSIColor(5)

	t := table.New().
		Headers("ID", "Name").
		BorderStyle(lipgloss.NewStyle().Foreground(accentColor)).
		StyleFunc(func(row, col int) lipgloss.Style {
			s := lipgloss.NewStyle()
			switch col {
			case 0:
				s = s.Width(3)
			case 1:
				s = s.Width(20)
			}
			switch row {
			case 0:
				s = s.Foreground(accentColor)
			}
			return s
		})

	for _, c := range clients {
		t.Row(strconv.Itoa(int(c.ID)), c.Name)
	}

	fmt.Println(t)

	return nil
}

func (cmd *ListCmd) Run(app *context.Context) error {
	db := app.GetDb()
	defer db.Close()
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
