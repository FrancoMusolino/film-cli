package printMovie

import (
	"fmt"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/program"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/pariz/gountries"
)

var (
	primary      = lipgloss.Color("#01FAC6")
	secondary    = lipgloss.Color("#40BDA3")
	focusedStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#01FAC6")).Bold(true)
)

type model struct {
	movie      movies.MovieDetail
	program    *program.Program
	stepNumber int
}

func InitialModel(movie movies.MovieDetail, program *program.Program) model {
	return model{movie: movie, program: program, stepNumber: 3}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.program.StepChan <- m.stepNumber - 1
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	queryCountry := gountries.New()

	primary := lipgloss.Color("#01FAC6")
	secondary := lipgloss.Color("#40BDA3")

	titleStyle := lipgloss.NewStyle().
		Foreground(primary).
		Bold(true).
		Underline(true).
		PaddingBottom(1)

	descStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#AAAAAA")).PaddingBottom(1)

	labelStyle := lipgloss.NewStyle().
		Foreground(secondary).
		Bold(true)

	valueStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF"))

	s := titleStyle.Render(m.movie.Title) + "\n"
	s += descStyle.Render(m.movie.Overview) + "\n"
	s += labelStyle.Render(fmt.Sprintf("%s: ", m.program.Translate("genre-label"))) + valueStyle.Render(strings.Join(genreNames(m.movie), ", ")) + "\n"

	if m.movie.Adult {
		s += labelStyle.Render(m.program.Translate("adults-only")) + "\n"
	}

	if len(m.movie.OriginCountry) > 0 {
		if originCountry, err := queryCountry.FindCountryByAlpha(m.movie.OriginCountry[0]); err == nil {
			country := originCountry.Name.Common
			if m.program.Lang == "es" {
				country = originCountry.Translations["SPA"].Common
			}

			s += labelStyle.Render(fmt.Sprintf("%s: ", m.program.Translate("origin-country-label")) + valueStyle.Render(country) + "\n")
		}
	}

	s += labelStyle.Render(fmt.Sprintf("%s: ", m.program.Translate("original-language-label"))) + valueStyle.Render(m.movie.OriginalLanguage) + "\n"
	s += labelStyle.Render(fmt.Sprintf("%s: ", m.program.Translate("votes-label"))) + valueStyle.Render(fmt.Sprintf("%d", m.movie.VoteCount)) + "\n"
	s += labelStyle.Render("Rating: ") + valueStyle.Render(fmt.Sprintf("%.2f", m.movie.VoteAverage)) + "\n\n"

	s += fmt.Sprintf(m.program.Translate("back", map[string]interface{}{"KeyStroke": fmt.Sprint(focusedStyle.Render("Esc"))}))
	s += "\n"

	return s
}

func genreNames(m movies.MovieDetail) []string {
	names := make([]string, 0, len(m.Genres))
	for _, g := range m.Genres {
		names = append(names, g.Name)
	}
	return names
}
