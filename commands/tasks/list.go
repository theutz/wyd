package tasks

import (
	"context"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/internal/log"
	"github.com/theutz/wyd/queries"
)

type ListCmd struct {
	q queries.Queries
	c context.Context
}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	l := log.Get()
	cmd.q = b.Queries
	cmd.c = b.Context

	tasks, err := cmd.q.ListTasks(cmd.c)
	if err != nil {
		l.Fatal(err)
	}
	l.Debug("tasks", "tasks", tasks)

	t := table.New().
		Headers("Task", "Project")

	for _, task := range tasks {
		t.Row(task.Name, task.ProjectName)
	}

	l.Printf("Tasks:\n%s", t)

	return nil
}
