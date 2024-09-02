package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/queries"
)

type ProjectsCmd struct {
	Add  AddProjectCmd  `cmd:"" help:"add a project"`
	List ListProjectCmd `cmd:"" help:"list all projects"`
}

func (cmd *ProjectsCmd) Run() error {
	return nil
}

type AddProjectCmd struct {
	Name   string `short:"n" help:"the name of the package"`
	Client string `short:"c" help:"the name of the client"`
	fields []huh.Field
	params queries.CreateProjectParams
}

func (cmd *AddProjectCmd) handleName() {
	if cmd.Name == "" {
		name := huh.NewInput().
			Title("Name").
			Value(&cmd.params.Name)
		log.Debug("name is empty. prompting for input")
		cmd.fields = append(cmd.fields, name)
	} else {
		cmd.params.Name = cmd.Name
	}
}

func (cmd *AddProjectCmd) handleClient(c *Context) {
	if cmd.Client == "" {
		log.Debug("client is empty. prompting for input.")
		log.Debug("loading clients")
		clients, err := c.queries.ListClients(c.dbCtx)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("loaded clients", "count", len(clients))

		log.Debug("converting clients to options")
		var options []huh.Option[int64]
		for _, c := range clients {
			o := huh.NewOption[int64](c.Name, c.ID)
			options = append(options, o)
		}
		log.Debug("converted clients to options", "options", options)

		client := huh.NewSelect[int64]().
			Title("Client").
			Options(options...).
			Value(&cmd.params.ClientID)
		cmd.fields = append(cmd.fields, client)
	} else {
		log.Debug("searching for client by name", "name", cmd.Client)
		client, err := c.queries.GetClientByName(c.dbCtx, cmd.Client)
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("found", "client", client)
		cmd.params.ClientID = client.ID
	}
}

func (cmd *AddProjectCmd) saveProject(c *Context) {
	log.Debug("creating project", "params", cmd.params)
	project, err := c.queries.CreateProject(c.dbCtx, cmd.params)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("project created", "project", project)
}

func (cmd *AddProjectCmd) runForm() {
	if len(cmd.fields) > 0 {
		form := huh.NewForm(
			huh.NewGroup(cmd.fields...),
		)
		err := form.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (cmd *AddProjectCmd) Run(c *Context) error {
	log.Debug("adding project")
	log.Debug("flag", "name", cmd.Name)
	log.Debug("flag", "client", cmd.Client)

	cmd.handleName()
	cmd.handleClient(c)
	cmd.runForm()
	cmd.saveProject(c)

	return nil
}

type ListProjectCmd struct{}

func (cmd *ListProjectCmd) Run(c *Context) error {
	log.Debug("listing projects")

	projects, err := c.queries.ListProjects(c.dbCtx)
	if err != nil {
		log.Fatal(err)
	}

	t := table.New().
		Headers("Project Name", "Client Name")

	for _, project := range projects {
		t.Row(project.Name, project.ClientName)
	}

	fmt.Println(t)

	return nil
}
