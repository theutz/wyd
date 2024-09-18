package app

type Cli struct {
	Config  ConfigCmd  `cmd:"" help:"view wyd configuration"`
	Client  ClientCmd  `cmd:"" help:"work with client list"`
	Project ProjectCmd `cmd:"" help:"work with project list"`
}
