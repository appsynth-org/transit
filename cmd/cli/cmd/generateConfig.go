package cmd

import (
	"fmt"

	"github.com/appsynth-org/transit/config"
	"github.com/spf13/cobra"
)

var generateConfigCmd = &cobra.Command{
	Use:   "generate",
	Short: "Command for generating config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🥷🏻 Generating config file...")
		config.GenerateConfig()
	},
}

func init() {
	configCmd.AddCommand(generateConfigCmd)
}
