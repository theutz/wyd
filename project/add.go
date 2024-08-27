package project

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/queries"
)

type AddCmd struct {
	Name   string `short:"n" help:"the name of the package"`
	Client string `short:"c" help:"the name of the client"`
	fields []huh.Field
	params queries.CreateProjectParams
	b      bindings.Bindings
}

func (cmd *AddCmd) handleName() {
	if cmd.Name == "" {
		name := huh.NewInput().
			Title("Name").
			Value(&cmd.params.Name)
		cmd.b.Logger.Debug("name is empty. prompting for input", "field", name)
		cmd.fields = append(cmd.fields, name)
	} else {
		cmd.params.Name = cmd.Name
	}
}

func (cmd *AddCmd) handleClient() {
	if cmd.Client == "" {
		cmd.b.Logger.Debug("client is empty. prompting for input.")
		cmd.b.Logger.Debug("loading clients")
		clients, err := cmd.b.Queries.ListClients(cmd.b.Context)
		if err != nil {
			cmd.b.Logger.Fatal(err)
		}
		cmd.b.Logger.Debug("loaded clients", "clients", clients)

		cmd.b.Logger.Debug("converting clients to options")
		var options []huh.Option[int64]
		for _, c := range clients {
			o := huh.NewOption[int64](c.Name, c.ID)
			options = append(options, o)
		}
		cmd.b.Logger.Debug("converted clients to options", "options", options)

		client := huh.NewSelect[int64]().
			Title("Client").
			Options(options...).
			Value(&cmd.params.ClientID)
		cmd.b.Logger.Debug("field", "client", client)
		cmd.fields = append(cmd.fields, client)
	} else {
		cmd.b.Logger.Debug("searching for client by name", "name", cmd.Client)
		client, err := cmd.b.Queries.GetClientByName(cmd.b.Context, cmd.Client)
		if err != nil {
			cmd.b.Logger.Fatal(err)
		}
		cmd.b.Logger.Debug("found", "client", client)
		cmd.params.ClientID = client.ID
	}
}

func (cmd *AddCmd) saveProject() {
	cmd.b.Logger.Debug("creating project", "params", cmd.params)
	project, err := cmd.b.Queries.CreateProject(cmd.b.Context, cmd.params)
	if err != nil {
		cmd.b.Logger.Fatal(err)
	}
	cmd.b.Logger.Info("project created", "project", project)
}

func (cmd *AddCmd) runForm() {
	if len(cmd.fields) > 0 {
		form := huh.NewForm(
			huh.NewGroup(cmd.fields...),
		)
		err := form.Run()
		if err != nil {
			cmd.b.Logger.Fatal(err)
		}
	}
}

func (cmd *AddCmd) Run(b bindings.Bindings) error {
	cmd.b = b
	b.Logger.Debug("adding project")
	b.Logger.Debug("flag", "name", cmd.Name)
	b.Logger.Debug("flag", "client", cmd.Client)

	cmd.handleName()
	cmd.handleClient()
	cmd.runForm()
	cmd.saveProject()

	return nil
}
