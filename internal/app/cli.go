package app

var cli struct {
	Config ConfigCmd `cmd:"" help:"view wyd configuration"`
	Client ClientCmd `cmd:"" help:"work with client list"`
}
