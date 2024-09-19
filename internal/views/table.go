package views

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Table(records Entries) string {
	r, ok := records.Get(0)
	if !ok {
		return getErrorText().Render("No records exist.")
	}

	headers := r.Keys()

	tbl := table.New().
		BorderStyle(getBorderStyle()).
		StyleFunc(func(row, col int) lipgloss.Style {
			style := lipgloss.NewStyle()

			switch col {
			case 0:
				style = style.Width(3)
			case 1:
				style = style.Width(20)
			}

			switch row {
			case 0:
				style = style.Foreground(getAccentColor())
			}

			return style
		}).
		Headers(headers...)

	it := records.Iterator()
	for it.Next() {
		entry := it.Value()
		row := entry.Values()
		tbl.Row(row...)
	}

	return tbl.Render()
}
