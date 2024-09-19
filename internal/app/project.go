package app

import (
	"context"
	"fmt"

	"github.com/theutz/wyd/internal/db/queries/clients"
	"github.com/theutz/wyd/internal/db/queries/projects"
)

type ProjectCmd struct {
	List   ProjectListCmd   `aliases:"show,ls"   cmd:"" help:"list all projects"`
	Add    ProjectAddCmd    `aliases:"create,a"  cmd:"" help:"add a new project"`
	Remove ProjectRemoveCmd `aliases:"delete,rm" cmd:"" help:"remove a project"`
}

type ProjectListCmd struct{}

func (cmd *ProjectListCmd) Run(ctx context.Context, p *projects.Queries) error {
	projects, err := p.All(ctx)
	if err != nil {
		return fmt.Errorf("loading all projects: %w", err)
	}

	fmt.Println(projects)

	return nil
}

type ProjectAddCmd struct {
	Name   string `help:"name of the project" required:"" short:"n"`
	Client string `help:"name of the client"  required:"" short:"c"`
}

func (cmd *ProjectAddCmd) Run(
	ctx context.Context,
	clientq *clients.Queries,
	projectq *projects.Queries,
) error {
	client, err := clientq.QueryByName(ctx, cmd.Client)
	if err != nil {
		return fmt.Errorf("querying client by name %s: %w", cmd.Client, err)
	}

	params := projects.CreateParams{
		Name:     cmd.Name,
		ClientID: client.ID,
	}

	project, err := projectq.Create(ctx, params)
	if err != nil {
		return fmt.Errorf("creating project: %w", err)
	}

	fmt.Println(project)

	return nil
}

type ProjectRemoveCmd struct {
	Name string `help:"the name of the project" required:"" short:"n"`
}

func (cmd *ProjectRemoveCmd) Run(ctx context.Context, q *projects.Queries) error {
	project, err := q.DeleteByName(ctx, cmd.Name)
	if err != nil {
		return fmt.Errorf("deleting project by name %s: %w", cmd.Name, err)
	}

	fmt.Println(project)

	return nil
}
