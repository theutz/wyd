package project

import (
	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/client"
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Name     string `gorm:"not null,unique"`
	ClientID int
	Client   client.Client
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

func (cmd *AddCmd) Run(log *clog.Logger, db *gorm.DB) error {
	var clients []client.Client

	if result := db.Find(&clients); result.Error != nil {
		log.Fatal(result.Error)
	} else {
		log.Debug("Clients found", "count", result.RowsAffected)
	}

	var options []huh.Option[int]
	for _, c := range clients {
		log.Debug("Client", "name", c.Name, "id", c.ID)
		options = append(options, huh.NewOption[int](c.Name, int(c.ID)))
	}

	var name string
	var clientId int
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Name").
				Value(&name),
			huh.NewSelect[int]().
				Title("Client").
				Options(options...).
				Value(&clientId),
		))

	err := form.Run()
	if err != nil {
		log.Fatal(err)
	}
	log.Debug(
		"Form",
		"name", name,
		"clientId", clientId,
	)

	db.Create(&Project{Name: name, ClientID: clientId})

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
