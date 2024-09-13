package clients

import (
	"github.com/theutz/wyd/internal/db/queries"
)

type ClientsCmd struct {
	Add    AddClientCmd     `cmd:"" help:"add a client"`
	List   ListClientsCmd   `cmd:"" default:"withargs" help:"list all clients"`
	Delete DeleteClientsCmd `cmd:"" help:"delete a client"`
}

type AddClientCmd struct {
	Name   string `short:"n" required:"" help:"the name of the client"`
	client queries.Client
}

type ListClientsCmd struct {
	clients []queries.Client
}

type DeleteClientsCmd struct {
	Name   string `short:"n" xor:"id" required:"" help:"the name of the client"`
	Id     int64  `short:"i" xor:"id" required:"" help:"the id of the client"`
	client queries.Client
}
