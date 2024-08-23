package project

import (
	"github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name string
}

type ProjectCmd struct {
	Add    AddCmd    `cmd:"" help:"add a project"`
	List   ListCmd   `cmd:"" help:"list all projects"`
	Remove RemoveCmd `cmd:"" help:"remove a project"`
}

func (cmd *ProjectCmd) Run() error {
	return nil
}

type AddCmd struct {
	Name string `short:"n" help:"the name of the package"`
}

func (cmd *AddCmd) Run(l *log.Logger, db *gorm.DB) error {
	db.Create(&Project{Name: "meyh"})
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
