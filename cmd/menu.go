package cmd

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/ui/multiInput"
	"github.com/FrancoMusolino/film-cli/cmd/ui/printMovie"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
}

type Options struct {
	MenuItem  *multiInput.Selection
	MovieItem *multiInput.Selection
}

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Navega sobre nuestro menú y explora las mejores películas",
	Long:  "Explora las mejores películas de la historia, las que se están reproduciendo ahora, las más populares y demás!! Ideal para un pasionado del séptimo arte",
	Run: func(cmd *cobra.Command, args []string) {
		var wg sync.WaitGroup
		stepChan := make(chan int, 10)

		var tprogram *tea.Program
		menu := menu.InitMenu()
		moviesService := movies.NewMoviesService()

		options := Options{
			MenuItem:  &multiInput.Selection{},
			MovieItem: &multiInput.Selection{},
		}

		wg.Add(1)
		go func() {
			defer wg.Done()

			for i := range stepChan {
				switch i {
				case 1:
					fmt.Println(options.MenuItem)
					tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Items, options.MenuItem, "Elige una opción de nuestro menú", 1, stepChan))
					if _, err := tprogram.Run(); err != nil {
						log.Fatal(err)
					}
				case 2:
					s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
					s.Start()
					movies := getMoviesBySelectedMenuKey(options.MenuItem.Choice, moviesService)
					menu.SetMenuMovies(movies)
					s.Stop()

					tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Movies, options.MovieItem, "Elige una película", 2, stepChan))
					if _, err := tprogram.Run(); err != nil {
						log.Fatal(err)
					}
				case 3:
					i, _ := strconv.Atoi(options.MovieItem.Choice)

					s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
					s.Start()
					fmt.Println()
					detail, _ := moviesService.GetMovieDetail(i)
					s.Stop()

					printMovie.Print(detail)

				default:
					close(stepChan)
				}

			}
		}()

		stepChan <- 1
		wg.Wait()
	},
}

func getMoviesBySelectedMenuKey(key string, moviesService movies.MoviesService) []movies.Movie {
	var movies []movies.Movie

	switch key {
	case "top-rated":
		movies = moviesService.GetTopRatedMovies()
	case "popular":
		movies = moviesService.GetPopularMovies()
	case "now-playing":
		movies = moviesService.GetNowPlayingMovies()
	case "upcoming":
		movies = moviesService.GetUpcomingMovies()
	}

	return movies[:6]
}
