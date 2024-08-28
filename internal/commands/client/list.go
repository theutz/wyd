package client

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	q := db.Query

	log.Debug("listing clients")
	clients, err := q.ListClients(ctx)
	if err != nil {
		log.Fatal(err)
	}
	t := table.New().
		Headers("Name")
	for _, client := range clients {
		t.Row(client.Name)
	}
	fmt.Println(t)
	return nil
}
