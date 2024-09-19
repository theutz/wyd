package views

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Record(record Entry) string {
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
		tbl.Row(iter.Key(), iter.Value())
	}

	return tbl.Render()
}
