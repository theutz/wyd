package views

import "github.com/charmbracelet/lipgloss"

func getAccentColor() lipgloss.ANSIColor {
	return lipgloss.ANSIColor(5)
}

func getErrorColor() lipgloss.ANSIColor {
	return lipgloss.ANSIColor(1)
}

func getBorderStyle() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(getAccentColor())
}

func getErrorText() lipgloss.Style {
	return lipgloss.NewStyle().
		Foreground(getErrorColor())
}
