package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// generateTemplateCmd represents the generateTemplate command
var generateTemplateCmd = &cobra.Command{
	Use:   "generate-template",
	Short: "generate a sample config template",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("generateTemplate called")
	},
}

func init() {
	configCmd.AddCommand(generateTemplateCmd)

	// generateTemplateCmd.Flags().String("output-file", "t", false, "Help message for toggle")
}
