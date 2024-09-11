package clients

import "github.com/theutz/wyd/internal/db/clients"

type ClientsCmd struct {
	Add  AddCmd  `cmd:"" help:"add a client"`
	List ListCmd `cmd:"" default:"withargs" help:"list all clients"`
}

type AddCmd struct {
	Name   string `arg:"" help:"the name of the client"`
	client clients.Client
}

type ListCmd struct {
	clients []clients.Client
}
