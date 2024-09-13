package projects

type ProjectsCmd struct {
	Add    AddProjectsCmd    `cmd:"" help:"add a project"`
	List   ListProjectsCmd   `cmd:"" default:"withargs" help:"list all projects"`
	Delete DeleteProjectsCmd `cmd:"" help:"delete a project"`
}

type AddProjectsCmd struct {
	Name string `short:"n" help:"the project name"`
}

type ListProjectsCmd struct{}

type DeleteProjectsCmd struct {
	Name string `short:"n" xor:"project" required:"" help:"the project name"`
	Id   int64  `short:"i" xor:"project" required:"" help:"the project id"`
}
