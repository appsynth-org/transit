/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"context"
	"fmt"

	"github.com/appsynth-org/transit/config"
	"github.com/appsynth-org/transit/reader"
	"github.com/appsynth-org/transit/writer"
	"github.com/spf13/cobra"
)

// var (
// 	output string
// )

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Read Google Sheet then generate the locale files",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("♻️  Generating locale files...")
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		ctx := context.Background()
		groups, err := reader.ReadGoogleSheet(config, ctx)
		if err != nil {
			fmt.Printf("Unable to read spreadsheet %v", err)
			return
		}
		writer.GenerateLocaleFiles(groups)

		fmt.Println("✅ Generated locale files successfully, Please check the output folder.")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	// generateCmd.Flags().StringVarP(&output, "output", "o", "all", "Platform to generate locale files (ios|android)")

}
