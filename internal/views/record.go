package views

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/charmbracelet/log"
	"github.com/emirpasic/gods/maps/linkedhashmap"
)

func Record(record *linkedhashmap.Map) string {
	logger := log.WithPrefix("record")

	tbl := table.New().
		BorderStyle(getBorderStyle()).
		StyleFunc(func(row, col int) lipgloss.Style {
			sty := lipgloss.NewStyle()

			switch col {
			case 0:
				sty = sty.Width(10)
			case 1:
				sty = sty.Width(20)
			}

			switch row {
			case 0:
				sty = sty.Foreground(getAccentColor())
			}

			return sty
		})

	iter := record.Iterator()
	for iter.Next() {
		key, ok := iter.Key().(string) //nolint:varnamelen
		if !ok {
			logger.Fatal("key is not a string")
		}

		value, ok := iter.Value().(string)
		if !ok {
			logger.Fatal("value is not a string")
		}

		tbl.Row(key, value)
	}

	return tbl.Render()
}
