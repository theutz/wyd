package project

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/queries"
)

type AddCmd struct {
	Name     string `short:"n" help:"the name of the package"`
	ClientId int64  `xor:"client" help:"the client id (can't be used with --client)"`
	Client   string `short:"c" xor:"client" help:"the name of the client (can't be used with --client-id)"`
}

func (cmd *AddCmd) Run(b bindings.Bindings) error {
	b.Logger.Debug("adding project")
	b.Logger.Debug("flag", "name", cmd.Name)
	b.Logger.Debug("flag", "clientId", cmd.ClientId)
	b.Logger.Debug("flag", "client", cmd.Client)

	fields := []huh.Field{}

	params := queries.CreateProjectParams{
		Name:     cmd.Name,
		ClientID: int64(cmd.ClientId),
	}

	if cmd.Name == "" {
		name := huh.NewInput().
			Title("Name").
			Value(&params.Name)
		b.Logger.Debug("name is empty. prompting for input", "field", name)
		fields = append(fields, name)
	}

	if cmd.ClientId == 0 && cmd.Client == "" {
		b.Logger.Debug("client is empty. prompting for input.")
		b.Logger.Debug("loading clients")
		clients, err := b.Queries.ListClients(b.Context)
		if err != nil {
			b.Logger.Fatal(err)
		}
		b.Logger.Debug("loaded clients", "clients", clients)

		b.Logger.Debug("converting clients to options")
		var options []huh.Option[int64]
		for _, c := range clients {
			o := huh.NewOption[int64](c.Name, c.ID)
			options = append(options, o)
		}
		b.Logger.Debug("converted clients to options", "options", options)

		client := huh.NewSelect[int64]().
			Title("Client").
			Options(options...).
			Value(&params.ClientID)
		b.Logger.Debug("field", "client", client)
		fields = append(fields, client)
	}

	if params.ClientID == 0 {
		b.Logger.Debug("searching for client by name", "name", cmd.Client)
		client, err := b.Queries.GetClientByName(b.Context, cmd.Client)
		if err != nil {
			b.Logger.Fatal(err)
		}
		b.Logger.Debug("found", "client", client)
		params.ClientID = client.ID
	}

	if len(fields) > 0 {
		form := huh.NewForm(
			huh.NewGroup(fields...),
		)
		err := form.Run()
		if err != nil {
			b.Logger.Fatal(err)
		}
	}

	b.Logger.Debug("creating project", "params", params)
	project, err := b.Queries.CreateProject(b.Context, params)
	if err != nil {
		b.Logger.Fatal(err)
	}
	b.Logger.Info("project created", "project", project)

	return nil
}
