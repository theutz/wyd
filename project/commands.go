package project

type ProjectCmd struct {
	Add  AddCmd  `cmd:"" help:"add a project"`
	List ListCmd `cmd:"" help:"list all projects"`
	// Remove RemoveCmd `cmd:"" help:"remove a project"`
}

func (cmd *ProjectCmd) Run() error {
	return nil
}

// type RemoveCmd struct{}
//
// func (cmd *RemoveCmd) Run() error {
// 	return nil
// }
