package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "starwars-cli",
	Short: "Star Wars CLI",
	Long:  `Star Wars CLI`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {

	// Commands registered in the APP
	rootCmd.AddCommand(filmsCmd)
	rootCmd.AddCommand(planetsCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
