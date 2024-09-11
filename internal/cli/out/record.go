package out

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

func Record(record map[string]string) string {
	t := table.New().
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			s := lipgloss.NewStyle()
			switch col {
			case 0:
				s = s.Width(10)
			case 1:
				s = s.Width(20)
			}
			switch row {
			case 0:
				s = s.Foreground(accentColor)
			}
			return s
		})

	for k, v := range record {
		t.Row(k, v)
	}

	return t.Render()
}
