package tasks

type TaskCmd struct {
	Add  AddCmd  `cmd:"" help:"add a new task"`
	List ListCmd `cmd:"" help:"list all tasks"`
}

func (cmd *TaskCmd) Run() error {
	return nil
}
