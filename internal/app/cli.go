package app

var cli struct {
	Config ConfigCmd `cmd:"" help:"view wyd configuration"`
}

type ConfigCmd struct {
	Show ShowCmd `cmd:"" default:"withargs" help:"print config to stdout"`
}

type ShowCmd struct{}

func (cmd *ShowCmd) Run(app Application) error {
	l := app.Logger().WithPrefix("config show")
	l.Print(app.Config())
	return nil
}
