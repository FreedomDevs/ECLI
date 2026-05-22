package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Wizard struct {
	step int

	textInput textinput.Model

	items  []string
	cursor int

	ProjectName string
	Choice      string
}

func NewWizard(items []string) Wizard {
	ti := textinput.New()
	ti.Placeholder = "project-name"
	ti.Focus()
	ti.CharLimit = 64

	return Wizard{
		step:      1,
		textInput: ti,
		items:     items,
	}
}

func (m Wizard) Init() tea.Cmd {
	return textinput.Blink
}

func (m Wizard) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch m.step {

	case 1:
		m.textInput, cmd = m.textInput.Update(msg)

		if k, ok := msg.(tea.KeyMsg); ok {
			switch k.Type {

			case tea.KeyEnter:
				val := m.textInput.Value()
				if val != "" {
					m.ProjectName = val
					m.step = 2
				}
				return m, nil

			case tea.KeyEsc, tea.KeyCtrlC:
				return m, tea.Quit
			}
		}

		return m, cmd

	case 2:
		switch msg := msg.(type) {

		case tea.KeyMsg:
			switch msg.Type {

			case tea.KeyUp:
				if m.cursor > 0 {
					m.cursor--
				}

			case tea.KeyDown:
				if m.cursor < len(m.items)-1 {
					m.cursor++
				}

			case tea.KeyEnter:
				m.Choice = m.items[m.cursor]
				return m, tea.Quit

			case tea.KeyEsc, tea.KeyCtrlC:
				return m, tea.Quit
			}
		}
	}

	return m, cmd
}

func (m Wizard) View() string {

	if m.step == 1 {

		header := TitleStyle.Render("📦 Create new project")
		label := LabelStyle.Render("Project name")

		box := InputBoxStyle.Render(m.textInput.View())

		footer := SubtleStyle.Render("Press Enter to continue • Esc to quit")

		return lipgloss.JoinVertical(
			lipgloss.Left,
			header,
			"",
			label,
			box,
			"",
			footer,
		)
	}

	header := TitleStyle.Render("🚀 Choose template")

	var items string

	for i, item := range m.items {

		cursor := " "
		if i == m.cursor {
			cursor = ">"
		}

		line := fmt.Sprintf(
			"%s %s",
			CursorStyle.Render(cursor),
			item,
		)

		if i == m.cursor {
			line = SelectedStyle.Render(line)
		} else {
			line = ItemStyle.Render(line)
		}

		items += line + "\n"
	}

	footer := SubtleStyle.Render("↑/↓ navigate • Enter select • Esc quit")

	return lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		"",
		items,
		footer,
	)
}
