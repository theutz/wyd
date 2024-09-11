package clients

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/theutz/wyd/internal/db/clients"
)

func RenderTable(header []string, rows [][]string) string {
	if len(rows) == 0 && len(header) == 0 {
		return ""
	}

	accentColor := lipgloss.ANSIColor(5)

	t := table.New().
		Rows(rows...).
		BorderStyle(lipgloss.NewStyle().Foreground(accentColor)).
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

func printClients(clients []clients.Client) error {
	if len(clients) < 1 {
		return errors.New("no clients found")
	}

	header := []string{"ID", "Name"}
	rows := [][]string{}

	for _, c := range clients {
		rows = append(rows, []string{strconv.Itoa(int(c.ID)), c.Name})
	}

	t := RenderTable(header, rows)
	fmt.Println(t)

	fmt.Println(t)

	return nil
}
