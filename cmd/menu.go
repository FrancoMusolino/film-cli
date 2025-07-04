package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FrancoMusolino/film-cli/cmd/flags"
	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/ui/multiInput"
	"github.com/FrancoMusolino/film-cli/cmd/ui/printMovie"
	"github.com/FrancoMusolino/film-cli/cmd/utils"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	var lang flags.Lang

	rootCmd.AddCommand(menuCmd)
	menuCmd.Flags().VarP(&lang, "lang", "l", "Idioma de la aplicación (es o en)")

	utils.RegisterStaticCompletions(menuCmd, "lang", flags.AllowedLangs)
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
		flagLang := flags.Lang(cmd.Flag("lang").Value.String())
		if flagLang == "" {
			flagLang = flags.Lang(flags.DefaultLang)
		}

		fmt.Println(flagLang)

		stepChan := make(chan int, 10)
		doneChan := make(chan bool)

		var tprogram *tea.Program
		menu := menu.InitMenu()
		moviesService := movies.NewMoviesService(flagLang)

		options := Options{
			MenuItem:  &multiInput.Selection{},
			MovieItem: &multiInput.Selection{},
		}

		go func() {
			for i := range stepChan {
				switch i {
				case 1:
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

					tprogram = tea.NewProgram(printMovie.InitialModel(detail, stepChan))
					if _, err := tprogram.Run(); err != nil {
						log.Fatal(err)
					}

				default:
					doneChan <- true
				}

			}
		}()

		stepChan <- 1
		<-doneChan
		close(stepChan)
		close(doneChan)
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
