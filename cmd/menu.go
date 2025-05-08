package cmd

import (
	"fmt"
	"log"
	"time"

	"github.com/FrancoMusolino/film-cli/cmd/menu"
	"github.com/FrancoMusolino/film-cli/cmd/movies"
	"github.com/FrancoMusolino/film-cli/cmd/ui/multiInput"
	"github.com/briandowns/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(menuCmd)
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
		moviesService := movies.NewMoviesService()

		options := Options{
			MenuItem: &multiInput.Selection{},
		}

		tprogram = tea.NewProgram(multiInput.InitialModelMulti(menu.Items, options.MenuItem, "Elige una opción de nuestro menú"))
		if _, err := tprogram.Run(); err != nil {
			log.Fatal(err)
		}

		item, err := menu.FindItemOnMenu(options.MenuItem.Choice)
		if err != nil {
			log.Fatal(err)
		}

		s := spinner.New(spinner.CharSets[37], 100*time.Millisecond) // Build our new spinner
		s.Start()                                                    // Start the spinner
		movies := item.GetMovies(moviesService)
		s.Stop()
		fmt.Println(movies)
	},
}
