package cmd

import (
	"github.com/spf13/cobra"
)

var planetsCmd = &cobra.Command{
	Use:   "planets",
	Short: "Print the list of planets ",
	Long:  `Print the list of planets `,
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	planetsCmd.AddCommand(planetListCmd)
}
