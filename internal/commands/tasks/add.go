package tasks

type AddCmd struct {
	Name    string `short:"n" help:"the name of the task"`
	Project string `short:"p" help:"the name of the project"`
}

func (cmd *AddCmd) handleName() {
}

func (cmd *AddCmd) Run() error {
	cmd.handleName()

	return nil
}
