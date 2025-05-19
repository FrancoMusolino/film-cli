package printMovie

import (
	"fmt"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/charmbracelet/lipgloss"
	"github.com/pariz/gountries"
)

var (
	primary   = lipgloss.Color("#01FAC6")
	secondary = lipgloss.Color("#40BDA3")
)

func Print(m movies.MovieDetail) {
	queryCountry := gountries.New()

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

	// Render
	fmt.Println(titleStyle.Render(m.Title))
	fmt.Println(descStyle.Render(m.Overview))
	fmt.Println(labelStyle.Render("Géneros: ") + valueStyle.Render(strings.Join(genreNames(m), ", ")))

	if m.Adult {
		fmt.Println(labelStyle.Render("Es para adultos"))
	}

	if originCountry, err := queryCountry.FindCountryByAlpha(m.OriginCountry[0]); err == nil {
		fmt.Println(labelStyle.Render("País de origen: ") + valueStyle.Render(originCountry.Translations["SPA"].Common))

	}

	fmt.Println(labelStyle.Render("Lenjuage original: ") + valueStyle.Render(m.OriginalLanguage))
	fmt.Println(labelStyle.Render("Votos: ") + valueStyle.Render(fmt.Sprintf("%d", m.VoteCount)))
	fmt.Println(labelStyle.Render("Rating: ") + valueStyle.Render(fmt.Sprintf("%.2f", m.VoteAverage)))
}

func genreNames(m movies.MovieDetail) []string {
	names := make([]string, 0, len(m.Genres))
	for _, g := range m.Genres {
		names = append(names, g.Name)
	}
	return names
}
