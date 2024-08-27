package client

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/bindings"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	l.Debug("listing clients")
	clients, err := b.Queries.ListClients(b.Context)
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
