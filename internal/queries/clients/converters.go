package clients

import (
	"strconv"

	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func (c Client) ToMap() *linkedhashmap.Map {
	m := linkedhashmap.New()
	m.Put("ID", strconv.FormatInt(c.ID, 10))
	m.Put("Name", c.Name)

	return m
}
