package client

import (
	"fmt"

	"github.com/theutz/wyd/bindings"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	b.Logger.Debug("listing clients")
	clients, err := b.Queries.ListClients(*&b.Context)
	if err != nil {
		b.Logger.Fatal(err)
	}
	fmt.Println(clients)
	return nil
}
