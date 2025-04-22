package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:   "filmy",
	Short: "Bienvenido a Filmy! Una CLI para navegar entre las mejores películas del momento",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
