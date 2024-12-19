package cmd

import (
	"github.com/spf13/cobra"
)

var filmsCmd = &cobra.Command{
	Use:   "films",
	Short: "Print the list of films ordered by release date",
	Long:  `Print the list of films ordered by release date`,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	filmsCmd.AddCommand(filmsListCmd)
}
