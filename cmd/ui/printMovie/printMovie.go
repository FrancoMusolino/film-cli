package printMovie

import (
	"fmt"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
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
	stepChan   chan int
	stepNumber int
}

func InitialModel(movie movies.MovieDetail, stepChan chan int) model {
	return model{movie: movie, stepChan: stepChan, stepNumber: 3}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			m.stepChan <- m.stepNumber - 1
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
	s += labelStyle.Render("Géneros: ") + valueStyle.Render(strings.Join(genreNames(m.movie), ", ")) + "\n"

	if m.movie.Adult {
		s += labelStyle.Render("Es para adultos") + "\n"
	}

	if len(m.movie.OriginCountry) > 0 {
		if originCountry, err := queryCountry.FindCountryByAlpha(m.movie.OriginCountry[0]); err == nil {
			s += labelStyle.Render("País de origen: ") + valueStyle.Render(originCountry.Translations["SPA"].Common) + "\n"
		}
	}

	s += labelStyle.Render("Lenjuage original: ") + valueStyle.Render(m.movie.OriginalLanguage) + "\n"
	s += labelStyle.Render("Votos: ") + valueStyle.Render(fmt.Sprintf("%d", m.movie.VoteCount)) + "\n"
	s += labelStyle.Render("Rating: ") + valueStyle.Render(fmt.Sprintf("%.2f", m.movie.VoteAverage)) + "\n\n"

	s += fmt.Sprintf("Persiona %s para volver atrás.\n", focusedStyle.Render("Esc"))
	return s
}

func genreNames(m movies.MovieDetail) []string {
	names := make([]string, 0, len(m.Genres))
	for _, g := range m.Genres {
		names = append(names, g.Name)
	}
	return names
}
