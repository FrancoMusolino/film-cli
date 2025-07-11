package menu

import (
	"fmt"
	"strings"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/nicksnyder/go-i18n/i18n"
)

type Item struct {
	Key, Name, Headers string
}

var MainMenuItemKeys = []string{"top-rated", "popular", "now-playing", "upcoming"}

type Menu struct {
	Items  []Item
	Movies []Item
}

func InitMenu(t i18n.TranslateFunc) *Menu {
	menu := &Menu{
		Items: []Item{
			{
				Key:     "top-rated",
				Name:    t("top-rated.title"),
				Headers: t("top-rated.desc"),
			},
			{
				Key:     "now-playing",
				Name:    t("now-playing.title"),
				Headers: t("now-playing.desc"),
			},
			{
				Key:     "popular",
				Name:    t("popular.title"),
				Headers: t("popular.desc"),
			},
			{
				Key:     "upcoming",
				Name:    t("upcoming.title"),
				Headers: t("upcoming.desc"),
			},
		},
	}

	return menu
}

func (m *Menu) SetMenuMovies(movies []movies.Movie) {
	var items []Item

	for _, m := range movies {
		items = append(items, Item{Key: fmt.Sprintf("%v", m.ID), Name: m.Title, Headers: formatMovieOverview(m.Overview, 20)})
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
