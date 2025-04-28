package cmd

import (
	"fmt"

	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/ui/multiInput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
}

type Program struct {
	moviesService    movies.MoviesService
	selectedMenuItem menu.Item
}

type Options struct {
	MenuItem *multiInput.Selection
}

var menuCmd = &cobra.Command{
	Use:   "menu",
	Short: "Navega sobre nuestro menú y explora las mejores películas",
	Long:  "Explora las mejores películas de la historia, las que se están reproduciendo ahora, las más populares y demás!! Ideal para un pasionado del séptimo arte",
	Run: func(cmd *cobra.Command, args []string) {
		var tprogram *tea.Program
		menu := menu.InitMenu()

		options := Options{
			MenuItem: &multiInput.Selection{},
		}

		movieService := movies.NewMoviesService()

		program := Program{
			moviesService: movieService,
		}

		tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Items, options.MenuItem, "Elige una opción de nuestro menú"))
		if _, err := tprogram.Run(); err != nil {
			panic(err)
		}
		program.selectedMenuItem = *options.MenuItem.Choice

		movies := program.selectedMenuItem.GetMovies(program.moviesService)

		fmt.Println(movies)
	},
}
