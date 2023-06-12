/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/appsynth-org/transit/config"
	"github.com/spf13/cobra"
)

var (
	generate bool
	list     bool
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Command for listing and generating config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		if !generate && !list {
			fmt.Println("Please provide -h for more information")
			return
		}

		if generate {
			fmt.Println("🥷🏻 Generating config file...")

			config.GenerateConfig()
		}

		if list {
			fmt.Println("🧐 Listing config file...")
			config, err := config.LoadConfig()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Printf("SERVICE_ACCOUNT_BASE64=%s\n", config.SERVICE_ACCOUNT_BASE64)
			fmt.Printf("GOOGLE_SHEET_ID=%s\n", config.GOOGLE_SHEET_ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.Flags().BoolVarP(&generate, "generate", "g", false, "Generate .env file")
	configCmd.Flags().BoolVarP(&list, "list", "l", false, "List config file")
}
