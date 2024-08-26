package client

import (
	"context"
	"fmt"

	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(log *clog.Logger, q *queries.Queries) error {
	log.Debug("listing clients")
	clients, err := q.ListClients(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(clients)
	return nil
}
