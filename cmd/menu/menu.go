package menu

import "github.com/FrancoMusolino/film-cli/cmd/movies"

type Item struct {
	Name, Headers string
	GetMovies     func(s movies.MoviesService) []movies.Movie
}

type Menu struct {
	Items []Item
}

func InitMenu() *Menu {
	menu := &Menu{
		[]Item{
			{
				Name:    "Mejor calificadas",
				Headers: "Obten las películas mejores calificadas",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetTopRatedMovies()
				},
			},
			{
				Name:    "Reproduciendo ahora",
				Headers: "Obten las películas que se están reproduciendo en las pantallas de los cines",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetNowPlayingMovies()
				},
			},
			{
				Name:    "Populares",
				Headers: "Obten las películas más populares del momento",
				GetMovies: func(s movies.MoviesService) []movies.Movie {
					return s.GetPopularMovies()
				},
			},
			{
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
