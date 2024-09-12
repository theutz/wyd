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
	Name   string `arg:"" help:"the name of the client"`
	client queries.Client
}

type ListCmd struct {
	clients []queries.Client
}

type DeleteCmd struct {
	Name   string `arg:"" help:"the name of the client"`
	client queries.Client
}
