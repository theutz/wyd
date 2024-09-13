package tasks

type TasksCmd struct {
	Add    AddTaskCmd    `cmd:"" help:"add a task"`
	List   ListTasksCmd  `cmd:"" default:"withargs" help:"list all tasks"`
	Delete DeleteTaskCmd `cmd:"" help:"delete a task"`
}

type AddTaskCmd struct {
	Name    string `short:"n" required:"" help:"the task name"`
	Project string `short:"p" required:"" help:"the task's project"`
}

type ListTasksCmd struct {
	Project string `short:"p" help:"filter by project"`
}

type DeleteTaskCmd struct {
	Name string `short:"n" required:"" xor:"task" help:"delete by name"`
	Id   int64  `short:"i" required:"" xor:"task" help:"delete by id"`
}
