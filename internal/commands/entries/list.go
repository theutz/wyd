package entries

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/theutz/wyd/internal/db"
)

type ListCmd struct{}

func (cmd *ListCmd) Run() error {
	q := db.Query
	ctx := db.Ctx

	entries, err := q.ListEntries(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Debug("entries", "entries", entries)

	t := table.New().
		Headers("Entry Name", "Task Name")

	for _, entry := range entries {
		t.Row(entry.Name, entry.TaskName)
	}

	fmt.Println(t)

	return nil
}
