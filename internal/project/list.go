package project

import (
	"fmt"

	"github.com/theutz/wyd/internal/bindings"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(b bindings.Bindings) error {
	b.Logger.Debug("listing projects")

	projects, err := b.Queries.ListProjects(*&b.Context)
	if err != nil {
		b.Logger.Fatal(err)
	}

	fmt.Println(projects)

	return nil
}
