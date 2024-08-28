package tasks

import (
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	q := db.Query
	ctx := db.Ctx

	tasks, err := q.ListTasks(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("tasks", "tasks", tasks)

	t := table.New().
		Headers("Task", "Project")

	for _, task := range tasks {
		t.Row(task.Name, task.ProjectName)
	}

	log.Printf("Tasks:\n%s", t)

	return nil
}
