package client

import "github.com/theutz/wyd/internal/log"

type ClientCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" help:"list clients"`
	Remove RemoveCmd `cmd:"" help:"remove a client"`
}

var l = log.Get()

func (cmd *ClientCmd) Run() error {
	return nil
}
