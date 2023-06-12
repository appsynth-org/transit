package cmd

import (
	"fmt"

	"github.com/appsynth-org/transit/config"
	"github.com/spf13/cobra"
)

var listConfigCmd = &cobra.Command{
	Use:   "list",
	Short: "Command for listing config file",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("🧐 Listing config file...")
		config, err := config.LoadConfig()
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("SERVICE_ACCOUNT_BASE64=%s\n", config.SERVICE_ACCOUNT_BASE64)
		fmt.Printf("GOOGLE_SHEET_ID=%s\n", config.GOOGLE_SHEET_ID)
	},
}

func init() {
	configCmd.AddCommand(listConfigCmd)
}
