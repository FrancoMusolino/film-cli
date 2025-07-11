package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/FrancoMusolino/film-cli/cmd/flags"
	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/program"
	"github.com/FrancoMusolino/film-cli/cmd/ui/multiInput"
	"github.com/FrancoMusolino/film-cli/cmd/ui/printMovie"
	"github.com/FrancoMusolino/film-cli/cmd/utils"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/nicksnyder/go-i18n/i18n"
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

		T, err := i18n.Tfunc(string(flagLang), flags.AllowedLangs...)
		if err != nil {
			log.Fatal("Cannot load transaltions")
		}

		moviesService := movies.NewMoviesService(flagLang)

		program := program.Program{
			Lang:          flagLang,
			Translate:     T,
			MoviesService: moviesService,
			StepChan:      make(chan int, 10),
			DoneChan:      make(chan bool),
		}

		var tprogram *tea.Program
		menu := menu.InitMenu(program.Translate)

		options := Options{
			MenuItem:  &multiInput.Selection{},
			MovieItem: &multiInput.Selection{},
		}

		go func() {
			for i := range program.StepChan {
				switch i {
				case 1:
					tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Items, options.MenuItem, program.Translate("choose-option"), &program, 1))
					if _, err := tprogram.Run(); err != nil {
						log.Fatal(err)
					}
				case 2:
					s := spinner.New(spinner.CharSets[37], 100*time.Millisecond)
					s.Start()
					movies := getMoviesBySelectedMenuKey(options.MenuItem.Choice, moviesService)
					menu.SetMenuMovies(movies)
					s.Stop()

					tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Movies, options.MovieItem, program.Translate("choose-movie"), &program, 2))
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

					tprogram = tea.NewProgram(printMovie.InitialModel(detail, &program))
					if _, err := tprogram.Run(); err != nil {
						log.Fatal(err)
					}

				default:
					program.DoneChan <- true
				}

			}
		}()

		program.StepChan <- 1
		<-program.DoneChan
		program.Terminate()
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
