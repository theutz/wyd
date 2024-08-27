package project

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/log"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	l := log.Get()
	q := db.Query
	ctx := db.Ctx

	l.Debug("listing projects")

	projects, err := q.ListProjects(ctx)
	if err != nil {
		l.Fatal(err)
	}

	t := table.New().
		Headers("Name", "Client")

	for _, project := range projects {
		t.Row(project.Name, project.ClientName)
	}

	fmt.Println(t)

	return nil
}
