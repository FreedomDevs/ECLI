package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Model struct {
	items  []string
	cursor int
	Choice string
}

func NewSelector(items []string) *Model {
	return &Model{
		items: items,
	}
}

func (m *Model) Init() tea.Cmd {
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {

		case tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit

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
		}
	}

	return m, nil
}

func (m *Model) View() string {

	s := TitleStyle.Render("Choose template") + "\n\n"

	for i, item := range m.items {

		cursor := " "
		style := ItemStyle

		if m.cursor == i {
			cursor = ">"
			style = SelectedStyle
		}

		s += fmt.Sprintf("%s %s\n", CursorStyle.Render(cursor), style.Render(item))
	}

	return s + "\n↑/↓ move • enter select • esc quit"
}
