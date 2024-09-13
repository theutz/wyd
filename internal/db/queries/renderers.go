package queries

import (
	"strconv"

	"github.com/theutz/wyd/internal/cli/out"
)

type Renderable interface {
	Render() string
}

func (c *Client) Render() string {
	id := strconv.Itoa(int(c.ID))
	client := map[string]string{
		"ID":   id,
		"Name": c.Name,
	}

	return out.Record(client)
}

type Clients []Client

func (c *Clients) Render() string {
	a := *c
	if len(a) < 1 {
		return "no clients found"
	}

	headers := []string{"ID", "Name"}
	rows := [][]string{}

	for _, c := range a {
		id := strconv.Itoa(int(c.ID))
		row := []string{id, c.Name}
		rows = append(rows, row)
	}

	return out.Table(headers, rows)
}

// TODO: Create renderer for project
// TODO: Create renderer for projects
// TODO: Create renderer for task
// TODO: Create renderer for tasks
