package app

type Cli struct {
	Config  ConfigCmd  `aliases:""           cmd:"" help:"view wyd configuration"`
	Client  ClientCmd  `aliases:"c,clients"  cmd:"" help:"work with client list"`
	Project ProjectCmd `aliases:"p,projects" cmd:"" help:"work with project list"`
}
