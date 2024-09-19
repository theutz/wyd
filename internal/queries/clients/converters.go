package clients

import (
	"strconv"

	"github.com/theutz/wyd/internal/views"
)

func (c Client) ToEntry() views.Entry {
	m := views.NewEntry()
	m.Put("ID", strconv.FormatInt(c.ID, 10))
	m.Put("Name", c.Name)

	return m
}

type Clients []Client

func (cs Clients) ToEntries() views.Entries {
	l := views.NewEntries()
	for _, c := range cs {
		l.Add(c.ToEntry())
	}

	return l
}
