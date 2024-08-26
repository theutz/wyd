package client

type ClientCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" help:"list clients"`
	Remove RemoveCmd `cmd:"" help:"remove a client"`
}

func (cmd *ClientCmd) Run() error {
	return nil
}
