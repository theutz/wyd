package project

import (
	"github.com/charmbracelet/huh"
	"github.com/theutz/wyd/internal/db/queries"
	"github.com/theutz/wyd/internal/log"
)

type AddCmd struct {
	Name   string `short:"n" help:"the name of the package"`
	Client string `short:"c" help:"the name of the client"`
	fields []huh.Field
	params queries.CreateProjectParams
}

func (cmd *AddCmd) handleName() {
	if cmd.Name == "" {
		name := huh.NewInput().
			Title("Name").
			Value(&cmd.params.Name)
		l.Debug("name is empty. prompting for input")
		cmd.fields = append(cmd.fields, name)
	} else {
		cmd.params.Name = cmd.Name
	}
}

func (cmd *AddCmd) handleClient() {
	if cmd.Client == "" {
		l.Debug("client is empty. prompting for input.")
		l.Debug("loading clients")
		clients, err := q.ListClients(ctx)
		if err != nil {
			l.Fatal(err)
		}
		l.Debug("loaded clients", "count", len(clients))

		l.Debug("converting clients to options")
		var options []huh.Option[int64]
		for _, c := range clients {
			o := huh.NewOption[int64](c.Name, c.ID)
			options = append(options, o)
		}
		l.Debug("converted clients to options", "options", options)

		client := huh.NewSelect[int64]().
			Title("Client").
			Options(options...).
			Value(&cmd.params.ClientID)
		cmd.fields = append(cmd.fields, client)
	} else {
		l.Debug("searching for client by name", "name", cmd.Client)
		client, err := q.GetClientByName(ctx, cmd.Client)
		if err != nil {
			l.Fatal(err)
		}
		l.Debug("found", "client", client)
		cmd.params.ClientID = client.ID
	}
}

func (cmd *AddCmd) saveProject() {
	l.Debug("creating project", "params", cmd.params)
	project, err := q.CreateProject(ctx, cmd.params)
	if err != nil {
		l.Fatal(err)
	}
	l.Info("project created", "project", project)
}

func (cmd *AddCmd) runForm() {
	if len(cmd.fields) > 0 {
		form := huh.NewForm(
			huh.NewGroup(cmd.fields...),
		)
		err := form.Run()
		if err != nil {
			l.Fatal(err)
		}
	}
}

func (cmd *AddCmd) Run() error {
	l := log.Get()

	l.Debug("adding project")
	l.Debug("flag", "name", cmd.Name)
	l.Debug("flag", "client", cmd.Client)

	cmd.handleName()
	cmd.handleClient()
	cmd.runForm()
	cmd.saveProject()

	return nil
}
