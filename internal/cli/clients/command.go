package clients

import "fmt"

type ListCmd struct{}

type ClientsCmd struct {
	List ListCmd `cmd:"" default:"withargs" help:"list all clients"`
}

func (cmd *ListCmd) Run() error {
	fmt.Println("list all clients")
	return nil
}
