/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	config_cmd "github.com/appsynth-org/transit/cmd/cli/cmd/config"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "transit",
	Short: "CMD tool for generating locale files from Google Sheet",
	Long: `Transit is a command line tool for generating locale files from Google Sheet.
	
Prerequisite:
- Run 'transit config generate' to generate .env file
- Replace generated .env file with your own values
- Run 'transit generate' to generate locale files and saved in output folder
	`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(config_cmd.ConfigCmd)
}
