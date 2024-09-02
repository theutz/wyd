package tasks

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

	tasks, err := q.ListTasks(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("tasks", "tasks", tasks)

	t := table.New().
		Headers("Task Name", "Project Name")

	for _, task := range tasks {
		t.Row(task.Name, task.ProjectName)
	}

	fmt.Println(t)

	return nil
}
