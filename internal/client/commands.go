package client

import (
	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"gorm.io/gorm"
)

type Client struct {
	gorm.Model
	Name string `gorm:"unique,not null"`
}

type ClientCmd struct {
	Add    AddCmd    `cmd:"" help:"add a client"`
	List   ListCmd   `cmd:"" help:"list clients"`
	Remove RemoveCmd `cmd:"" help:"remove a client"`
}

func (cmd *ClientCmd) Run() error {
	return nil
}

type AddCmd struct{}

func (cmd *AddCmd) Run(log *clog.Logger, db *gorm.DB) error {
	var name string

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("Name").Value(&name),
		),
	)

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("Input", "name", name)

	db.Create(&Client{Name: name})
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
