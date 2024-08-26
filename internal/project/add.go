package project

import (
	"context"
	"database/sql"
	"errors"
	"strconv"

	"github.com/charmbracelet/huh"
	clog "github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/queries"
)

type clientId string

func validateClientId[T interface{ ~string }](c T) error {
	s, ok := any(c).(string)
	if !ok {
		return errors.New("client id is not a valid string")
	}
	if _, err := strconv.Atoi(s); err != nil {
		return errors.New("client id must be an integer")
	}
	return nil
}

func (c *clientId) Validate() error {
	return validateClientId(*c)
}

type AddCmd struct {
	Name     string   `short:"n" help:"the name of the package"`
	ClientId clientId `short:"c" help:"the client id"`
}

func (cmd *AddCmd) Run(log *clog.Logger, q *queries.Queries) error {
	log.Debug("adding project")
	log.Debug("flag", "name", cmd.Name)
	log.Debug("flag", "clientId", cmd.ClientId)

	p := queries.CreateProjectParams{
		Name: cmd.Name,
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

	if cmd.ClientId == "" {
		var clientId string
		err := huh.NewInput().
			Title("Client ID").
			Validate(validateClientId).
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
		p.ClientID = sql.NullInt64{Int64: val}
	}

	project, err := q.CreateProject(context.Background(), p)
	if err != nil {
		log.Fatal(err)
	}

	log.Info("project created", "name", project.Name, "client id", project.ClientID)

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
