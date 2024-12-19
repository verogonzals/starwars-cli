package cmd

import (
	platform "star-wars-cli/platform/films"

	"github.com/spf13/cobra"
)

var filmsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of films ordered by release date",
	Long:  `Print the list of films ordered by release date`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := platform.FilmsListByReleaseDate()
		return err
	},
}






