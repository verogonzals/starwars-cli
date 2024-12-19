package cmd

import (
	platform "star-wars-cli/platform/planets"

	"github.com/spf13/cobra"
)

var planetListCmd = &cobra.Command{
	Use:   "list",
	Short: "Print the list of films ordered by release date",
	Long:  `Print the list of films ordered by release date`,
	RunE: func(cmd *cobra.Command, args []string) error {
		err := platform.GetPlanetsList()
		return err
	},
}


func init(){}






