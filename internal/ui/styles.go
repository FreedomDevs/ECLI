package ui

import "github.com/charmbracelet/lipgloss"

var (
	TitleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	SubtleStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))

	LabelStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("39"))

	InputBoxStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("63")).
			Padding(0, 1)

	CursorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212")).
			Bold(true)

	ItemStyle = lipgloss.NewStyle().
			PaddingLeft(2)

	SelectedStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("46")).
			Bold(true)
)
