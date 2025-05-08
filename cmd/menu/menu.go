package menu

import (
	"errors"
	"slices"

	"github.com/FrancoMusolino/film-cli/cmd/movies"
)

type Item struct {
	Key, Name, Headers string
	GetMovies          func(s movies.MoviesService) []movies.Movie
}

var AllowedItemKeys = []string{"top-rated", "popular", "now-playing", "upcoming"}

func isValidItemKey(key string) bool {
	return slices.Contains(AllowedItemKeys, key)
}

type Menu struct {
	Items []Item
}

func InitMenu() *Menu {
	menu := &Menu{
		[]Item{
			{
				Key:     "top-rated",
				Name:    "Mejor calificadas",
				Headers: "Obten las películas mejores calificadas",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetTopRatedMovies()
				},
			},
			{
				Key:     "now-playing",
				Name:    "Reproduciendo ahora",
				Headers: "Obten las películas que se están reproduciendo en las pantallas de los cines",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetNowPlayingMovies()
				},
			},
			{
				Key:     "popular",
				Name:    "Populares",
				Headers: "Obten las películas más populares del momento",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetPopularMovies()
				},
			},
			{
				Key:     "upcoming",
				Name:    "Próximamente",
				Headers: "Obten las películas que estarán en cartelera en los próximos días",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetUpcomingMovies()
				},
			},
		},
	}

	return menu
}

func (m *Menu) FindItemOnMenu(key string) (*Item, error) {
	if !isValidItemKey(key) {
		return nil, errors.New("invalid menu item key")
	}

	for _, item := range m.Items {
		if item.Key == key {
			return &item, nil
		}
	}

	return nil, errors.New("key valid, but not added to menu list")
}
