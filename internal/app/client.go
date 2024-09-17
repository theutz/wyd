package app

type ClientCmd struct {
	List ClientListCmd `cmd:"" help:"list all clients"`
}

type ClientListCmd struct{}

func (cmd *ClientListCmd) Run() error {
	return nil
}
