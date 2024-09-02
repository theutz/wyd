package project

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	q := db.Query
	ctx := db.Ctx

	log.Debug("listing projects")

	projects, err := q.ListProjects(ctx)
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
