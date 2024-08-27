package tasks

import (
	"context"
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/bindings"
	"github.com/theutz/wyd/queries"
)

type ListCmd struct {
	l log.Logger
	q queries.Queries
	c context.Context
}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	cmd.l = *b.Logger
	cmd.q = b.Queries
	cmd.c = b.Context

	tasks, err := cmd.q.ListTasks(cmd.c)
	if err != nil {
		cmd.l.Fatal(err)
	}
	cmd.l.Debug("tasks", "tasks", tasks)

	t := table.New().
		Headers("Task", "Project")

	for _, task := range tasks {
		t.Row(task.Name, task.ProjectName)
	}

	fmt.Println(t)

	return nil
}
