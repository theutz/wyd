package out

import (
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
)

var (
	accentColor = lipgloss.ANSIColor(5)
	borderStyle = lipgloss.NewStyle().Foreground(accentColor)
)

func Table(header []string, rows [][]string) string {
	if len(rows) == 0 && len(header) == 0 {
		return ""
	}

	t := table.New().
		Rows(rows...).
		BorderStyle(borderStyle).
		StyleFunc(func(row, col int) lipgloss.Style {
			s := lipgloss.NewStyle()
			switch col {
			case 0:
				s = s.Width(3)
			case 1:
				s = s.Width(20)
			}
			switch row {
			case 0:
				s = s.Foreground(accentColor)
			}
			return s
		})

	if len(header) > 0 {
		t = t.Headers(header...)
	}

	return t.Render()
}
