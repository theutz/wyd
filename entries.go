package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
)

type EntriesCmd struct {
	Add  AddEntriesCmd  `cmd:"" help:"add a new entry"`
	List ListEntriesCmd `cmd:"" help:"list all entries"`
}

func (cmd *EntriesCmd) Run() error {
	return nil
}

type ListEntriesCmd struct{}

func (cmd *ListEntriesCmd) Run(c *Context) error {
	entries, err := c.queries.ListEntries(c.dbCtx)
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

type AddEntriesCmd struct {
	Name    string `short:"n" help:"the name of the entry"`
	Project string `short:"t" help:"the name of the task"`
}

func (cmd *AddEntriesCmd) Run() error {
	return nil
}
