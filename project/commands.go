package project

import (
	"database/sql"

	"github.com/charmbracelet/log"
)

type ProjectCmd struct {
	Add    AddCmd    `cmd:"" help:"add a project"`
	List   ListCmd   `cmd:"" help:"list all projects"`
	Remove RemoveCmd `cmd:"" help:"remove a project"`
}

func (cmd *ProjectCmd) Run() error {
	// db, err := db.Open()
	return nil
}

type AddCmd struct{}

func (cmd *AddCmd) Run(l *log.Logger, db *sql.DB) error {
	return nil
}

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	return nil
}

type RemoveCmd struct{}

func (cmd *RemoveCmd) Run() error {
	return nil
}
