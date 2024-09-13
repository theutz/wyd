package projects

type ProjectsCmd struct {
	Add    AddProjectsCmd    `cmd:"" aliases:"a" help:"add a project"`
	List   ListProjectsCmd   `cmd:"" default:"withargs" help:"list all projects"`
	Delete DeleteProjectsCmd `cmd:"" aliases:"d" help:"delete a project"`
}

type AddProjectsCmd struct {
	Name   string `short:"n" help:"the project name"`
	Client string `short:"c" help:"the project's client"`
}

type ListProjectsCmd struct {
	Client string `short:"c" help:"filter by client"`
}

type DeleteProjectsCmd struct {
	Name string `short:"n" xor:"project" required:"" help:"the project name"`
	Id   int64  `short:"i" xor:"project" required:"" help:"the project id"`
}
