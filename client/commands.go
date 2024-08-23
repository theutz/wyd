package client

type ClientCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" help:"list clients"`
	Remove RemoveCmd `cmd:"" help:"remove a client"`
}

func (cmd *ClientCmd) Run() error {
	return nil
}

type AddCmd struct{}

func (cmd *AddCmd) Run() error {
	return nil
}

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	return nil
}

type RemoveCmd struct{}

func (cmd *RemoveCmd) Run() error {
	return nil
}
