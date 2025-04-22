package multiInput

import (
	"github.com/FrancoMusolino/film-cli/cmd/menu"
	tea "github.com/charmbracelet/bubbletea"
)

type Selection struct {
	Choice string
}

type model struct {
	choices  []menu.Item
	cursor   int
	selected map[int]struct{}
	choice   *Selection
	header   string
}

func (m model) Init() tea.Cmd {
	return nil
}

func InitialModelMulti(choices []menu.Item, selection *Selection, header string) model {
	return model{
		choices:  choices,
		selected: make(map[int]struct{}),
		choice:   selection,
		// header:   titleStyle.Render(header),
	}
}
