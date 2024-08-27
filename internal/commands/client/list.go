package client

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/log"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	l := log.Get()
	q := db.Query

	l.Debug("listing clients")
	clients, err := q.ListClients(ctx)
	if err != nil {
		l.Fatal(err)
	}
	t := table.New().
		Headers("Name")
	for _, client := range clients {
		t.Row(client.Name)
	}
	fmt.Println(t)
	return nil
}
