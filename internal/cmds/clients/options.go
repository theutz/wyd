package clients

type ClientsCmd struct {
	Add    AddClientCmd     `cmd:"" aliases:"a" help:"add a client"`
	List   ListClientsCmd   `cmd:"" default:"withargs" help:"list all clients"`
	Delete DeleteClientsCmd `cmd:"" aliases:"d" help:"delete a client"`
}

type AddClientCmd struct {
	Name string `short:"n" required:"" help:"the name of the client"`
}

type ListClientsCmd struct{}

type DeleteClientsCmd struct {
	Name string `short:"n" xor:"id" required:"" help:"the name of the client"`
	Id   int64  `short:"i" xor:"id" required:"" help:"the id of the client"`
}
