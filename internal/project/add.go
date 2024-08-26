package project

import (
	"context"
	"strconv"

	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type AddCmd struct {
	Name     string `short:"n" help:"the name of the package"`
	ClientId int    `xor:"client" help:"the client id (can't be used with --client)"`
	Client   string `short:"c" xor:"client" help:"the name of the client (can't be used with --client-id)"`
}

func (cmd *AddCmd) Run(log *clog.Logger, q *queries.Queries) error {
	log.Debug("adding project")
	log.Debug("flag", "name", cmd.Name)
	log.Debug("flag", "clientId", cmd.ClientId)
	log.Debug("flag", "client", cmd.Client)

	p := queries.CreateProjectParams{
		Name:     cmd.Name,
		ClientID: int64(cmd.ClientId),
	}

	if cmd.Name == "" {
		err := huh.NewInput().
			Title("Name").
			Value(&p.Name).
			Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("input", "name", p.Name)
	}

	if cmd.ClientId == 0 {
		var clientId string
		err := huh.NewInput().
			Title("Client ID").
			Value(&clientId).
			Run()
		log.Debug("input", "clientId", clientId)
		if err != nil {
			log.Fatal(err)
		}
		val, err := strconv.ParseInt(clientId, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		p.ClientID = val
	}

	log.Debug("creating project", "params", p)
	project, err := q.CreateProject(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("project created", "project", project)

	// var clients []client.Client
	//
	// if result := db.Find(&clients); result.Error != nil {
	// 	log.Fatal(result.Error)
	// } else {
	// 	log.Debug("Clients found", "count", result.RowsAffected)
	// }
	//
	// var options []huh.Option[int]
	// for _, c := range clients {
	// 	log.Debug("Client", "name", c.Name, "id", c.ID)
	// 	options = append(options, huh.NewOption[int](c.Name, int(c.ID)))
	// }
	//
	// var name string
	// var clientId int
	// form := huh.NewForm(
	// 	huh.NewGroup(
	// 		huh.NewInput().
	// 			Title("Name").
	// 			Value(&name),
	// 		huh.NewSelect[int]().
	// 			Title("Client").
	// 			Options(options...).
	// 			Value(&clientId),
	// 	))
	//
	// err := form.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Debug(
	// 	"Form",
	// 	"name", name,
	// 	"clientId", clientId,
	// )

	// db.Create(&Project{Name: name, ClientID: clientId})

	return nil
}
