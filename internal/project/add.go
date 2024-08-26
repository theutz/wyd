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

	params := queries.CreateProjectParams{
		Name:     cmd.Name,
		ClientID: int64(cmd.ClientId),
	}

	if cmd.Name == "" {
		err := huh.NewInput().
			Title("Name").
			Value(&params.Name).
			Run()
		if err != nil {
			log.Fatal(err)
		}
		log.Debug("input", "name", params.Name)
	}

	if params.ClientID == 0 {
		var clientIdStr string
		var clientId int64

		if cmd.Client == "" {
			clients, err := q.ListClients(context.Background())
			var options []huh.Option[int64]
			for _, c := range clients {
				o := huh.NewOption[int64](c.Name, c.ID)
				options = append(options, o)
			}
			err = huh.NewSelect[int64]().
				Options(options...).
				Value(&clientId).
				Run()
				// err := huh.NewInput().
				// 	Title("Client ID").
				// 	Value(&clientIdStr).
				// 	Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Debug("input", "clientId", clientIdStr)
		} else {
			var err error
			clientId, err = strconv.ParseInt(clientIdStr, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
		}
		params.ClientID = clientId
	}

	log.Debug("creating project", "params", params)
	project, err := q.CreateProject(context.Background(), params)
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
