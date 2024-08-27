package project

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/bindings"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	b.Logger.Debug("listing projects")

	projects, err := b.Queries.ListProjects(b.Context)
	if err != nil {
		b.Logger.Fatal(err)
	}

	t := table.New().
		Headers("Name", "Client")

	for _, project := range projects {
		t.Row(project.Name, project.ClientName)
	}

	fmt.Println(t)

	return nil
}
