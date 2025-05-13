package menu

import (
	"fmt"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
)

type Item struct {
	Key, Name, Headers string
}

var MainMenuItemKeys = []string{"top-rated", "popular", "now-playing", "upcoming"}

type Menu struct {
	Items  []Item
	Movies []Item
}

func InitMenu() *Menu {
	menu := &Menu{
		Items: []Item{
			{
				Key:     "top-rated",
				Name:    "Mejor calificadas",
				Headers: "Obten las películas mejores calificadas",
			},
			{
				Key:     "now-playing",
				Name:    "Reproduciendo ahora",
				Headers: "Obten las películas que se están reproduciendo en las pantallas de los cines",
			},
			{
				Key:     "popular",
				Name:    "Populares",
				Headers: "Obten las películas más populares del momento",
			},
			{
				Key:     "upcoming",
				Name:    "Próximamente",
				Headers: "Obten las películas que estarán en cartelera en los próximos días",
			},
		},
	}

	return menu
}

func (m *Menu) SetMenuMovies(movies []movies.Movie) {
	var items []Item

	for _, m := range movies {
		items = append(items, Item{Key: m.OriginalTitle, Name: m.Title, Headers: formatMovieOverview(m.Overview, 20)})
	}

	m.Movies = items
}

func formatMovieOverview(overview string, maxWords int) string {
	s := strings.Split(overview, " ")

	if len(s) <= maxWords {
		return strings.Join(s, " ")
	}

	return fmt.Sprintf("%s...", strings.Join(s[:maxWords], " "))
}
