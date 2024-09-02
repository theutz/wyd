package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
)

type TasksCmd struct {
	Add  AddTaskCmd   `cmd:"" help:"add a new task"`
	List ListTasksCmd `cmd:"" help:"list all tasks"`
}

func (cmd *TasksCmd) Run() error {
	return nil
}

type ListTasksCmd struct{}

func (cmd *ListTasksCmd) Run(c *Context) error {
	tasks, err := c.queries.ListTasks(c.dbCtx)
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

type AddTaskCmd struct {
	Name    string `short:"n" help:"the name of the task"`
	Project string `short:"p" help:"the name of the project"`
}

func (cmd *AddTaskCmd) handleName() {
}

func (cmd *AddTaskCmd) Run() error {
	cmd.handleName()

	return nil
}
