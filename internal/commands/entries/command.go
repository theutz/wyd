package entries

type EntryCmd struct {
	Add  AddCmd  `cmd:"" help:"add a new entry"`
	List ListCmd `cmd:"" help:"list all entries"`
}

func (cmd *EntryCmd) Run() error {
	return nil
}
