package project

import (
	"context"

	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type AddCmd struct {
	Name     string `short:"n" help:"the name of the package"`
	ClientId int    `xor:"client" help:"the client id (can't be used with --client)"`
	Client   string `short:"c" xor:"client" help:"the name of the client (can't be used with --client-id)"`
}

func (cmd *AddCmd) Run(log *clog.Logger, q *queries.Queries, ctx *context.Context) error {
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
		var clientId int64

		if cmd.Client == "" {
			log.Debug("loading clients")
			clients, err := q.ListClients(*ctx)
			if err != nil {
				log.Fatal(err)
			}
			log.Debug("loaded clients", "clients", clients)

			log.Debug("converting clients to options")
			var options []huh.Option[int64]
			for _, c := range clients {
				o := huh.NewOption[int64](c.Name, c.ID)
				options = append(options, o)
			}
			log.Debug("converted clients to options", "options", options)

			log.Debug("prompting for client")
			err = huh.NewSelect[int64]().
				Options(options...).
				Value(&clientId).
				Run()
			if err != nil {
				log.Fatal(err)
			}
			log.Debug("input", "clientId", clientId)
		} else {
			log.Debug("searching for client by name", "name", cmd.Client)
			client, err := q.GetClientByName(*ctx, cmd.Client)
			if err != nil {
				log.Fatal(err)
			}
			log.Debug("found", "client", client)
			clientId = client.ID
		}
		params.ClientID = clientId
	}

	log.Debug("creating project", "params", params)
	project, err := q.CreateProject(*ctx, params)
	if err != nil {
		log.Fatal(err)
	}
	log.Info("project created", "project", project)

	return nil
}
