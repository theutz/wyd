package tasks

import (
	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/internal/db"
	"github.com/theutz/wyd/internal/log"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	l := log.Get()
	q := db.Query
	ctx := db.Ctx

	tasks, err := q.ListTasks(ctx)
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
