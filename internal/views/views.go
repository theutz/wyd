package views

import (
	dll "github.com/ugurcsen/gods-generic/lists/doublylinkedlist"
	lhm "github.com/ugurcsen/gods-generic/maps/linkedhashmap"
)

type (
	Entry   = *lhm.Map[string, string]
	Entries = *dll.List[Entry]
)

func NewEntry() Entry {
	return lhm.New[string, string]()
}

func NewEntries() Entries {
	return dll.New[Entry]()
}
