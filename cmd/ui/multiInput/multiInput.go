package multiInput

import (
	"fmt"

	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/program"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle          = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
	titleStyle            = lipgloss.NewStyle().Background(lipgloss.Color("#01FAC6")).Foreground(lipgloss.Color("#030303")).Bold(true).Padding(0, 1, 0)
	descStyle             = lipgloss.NewStyle().Foreground(lipgloss.Color("#40BDA3"))
	selectedItemStyle     = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170")).Bold(true)
	selectedItemDescStyle = lipgloss.NewStyle().PaddingLeft(1).Foreground(lipgloss.Color("170"))
)

type Selection struct {
	Choice string
}

func (s *Selection) Update(v string) {
	s.Choice = v
}

type model struct {
	choices    []menu.Item
	cursor     int
	choice     *Selection
	header     string
	program    *program.Program
	stepNumber int
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelMulti(choices []menu.Item, selection *Selection, header string, program *program.Program, stepNumber int) model {
	var cursor int
	for i, choice := range choices {
		if choice.Key == selection.Choice {
			cursor = i
			break
		}
	}

	return model{
		cursor:     cursor,
		choices:    choices,
		choice:     selection,
		header:     titleStyle.Render(header),
		program:    program,
		stepNumber: stepNumber,
	}
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			m.program.StepChan <- 0
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
			m.program.StepChan <- m.stepNumber + 1
			return m, tea.Quit

		case "esc":
			m.program.StepChan <- m.stepNumber - 1
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
			choice.Headers = selectedItemDescStyle.Render(choice.Headers)
		}

		name := focusedStyle.Render(choice.Name)
		desc := descStyle.Render(choice.Headers)

		s += fmt.Sprintf("%s %s\n%s\n\n", cursor, name, desc)
	}

	s += fmt.Sprintf(m.program.Translate("foward", map[string]interface{}{"KeyStroke": fmt.Sprint(focusedStyle.Render("Enter"))}))
	s += "\n\n"
	s += fmt.Sprintf(m.program.Translate("back", map[string]interface{}{"KeyStroke": fmt.Sprint(focusedStyle.Render("Esc"))}))
	s += "\n"

	return s
}
