package project

import (
	"context"
	"fmt"

	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type ListCmd struct{}

func (cmd *ListCmd) Run(log *clog.Logger, q *queries.Queries, ctx *context.Context) error {
	log.Debug("listing projects")

	projects, err := q.ListProjects(*ctx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(projects)

	return nil
}
