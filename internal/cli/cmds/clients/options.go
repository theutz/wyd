package clients

import (
	"github.com/theutz/wyd/internal/db/queries"
)

type ClientsCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" default:"withargs" help:"list all clients"`
	Delete DeleteCmd `cmd:"" help:"delete a client"`
}

type AddCmd struct {
	Name   string `short:"n" required:"" help:"the name of the client"`
	client queries.Client
}

type ListCmd struct {
	clients []queries.Client
}

type DeleteCmd struct {
	Name   string `short:"n" xor:"id" required:"" help:"the name of the client"`
	Id     int64  `short:"i" xor:"id" required:"" help:"the id of the client"`
	client queries.Client
}
