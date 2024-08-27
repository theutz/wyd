package tasks

type TaskCmd struct {
	List ListCmd `cmd:"" help:"list all tasks"`
}

func (cmd *TaskCmd) Run() error {
	return nil
}
