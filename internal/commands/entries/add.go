package entries

type AddCmd struct {
	Name    string `short:"n" help:"the name of the entry"`
	Project string `short:"t" help:"the name of the task"`
}

func (cmd *AddCmd) handleName() {
}

func (cmd *AddCmd) Run() error {
	cmd.handleName()

	return nil
}
