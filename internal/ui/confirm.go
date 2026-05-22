package ui

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type Confirm struct {
	Message string

	options []string
	cursor  int

	Result bool
}

func NewConfirm(message string) *Confirm {
	return &Confirm{
		Message: message,
		options: []string{"Yes", "No"},
		cursor:  0,
	}
}

func (m *Confirm) Init() tea.Cmd {
	return nil
}

func (m *Confirm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	switch msg := msg.(type) {

	case tea.KeyMsg:

		switch msg.Type {

		case tea.KeyLeft, tea.KeyUp:
			if m.cursor > 0 {
				m.cursor--
			}

		case tea.KeyRight, tea.KeyDown:
			if m.cursor < len(m.options)-1 {
				m.cursor++
			}

		case tea.KeyEnter:
			m.Result = (m.cursor == 0) // Yes
			return m, tea.Quit

		case tea.KeyEsc, tea.KeyCtrlC:
			m.Result = false
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m *Confirm) View() string {

	s := TitleStyle.Render(m.Message) + "\n\n"

	for i, opt := range m.options {

		cursor := " "
		if i == m.cursor {
			cursor = ">"
			opt = SelectedStyle.Render(opt)
		} else {
			opt = ItemStyle.Render(opt)
		}

		s += fmt.Sprintf("%s %s\n", CursorStyle.Render(cursor), opt)
	}

	return s + "\n←/→ or ↑/↓ • enter confirm • esc cancel"
}
