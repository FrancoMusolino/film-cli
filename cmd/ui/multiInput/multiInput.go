package multiInput

import (
	"fmt"

	"github.com/FrancoMusolino/film-cli/cmd/menu"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle      = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	titleStyle        = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
)

type Selection struct {
	Choice string
}

func (s *Selection) Update(v string) {
	s.Choice = v
}

type model struct {
	choices []menu.Item
	cursor  int
	choice  *Selection
	header  string
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelMulti(choices []menu.Item, selection *Selection, header string) model {
	return model{
		choices: choices,
		choice:  selection,
		header:  titleStyle.Render(header),
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit

		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}

		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}

		case "enter", " ":
			m.choice.Update(m.choices[m.cursor].Key)
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := m.header + "\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if i == m.cursor {
			cursor = focusedStyle.Render(">")
			choice.Name = selectedItemStyle.Render(choice.Name)
		}

		name := focusedStyle.Render(choice.Name)

		s += fmt.Sprintf("%s %s\n\n", cursor, name)
	}

	s += fmt.Sprintf("Persiona %s para confirmar la selección.\n\n", focusedStyle.Render("Enter"))
	return s
}
