package cmd

import (
	"alertmanager/config"
	"fmt"

	"github.com/spf13/cobra"
)

// generateTemplateCmd represents the generateTemplate command
var generateTemplateCmd = &cobra.Command{
	Use:   "generate-template",
	Short: "generate a sample config template",

	Run: generateTemplateCmdRun,
}

func generateTemplateCmdRun(cmd *cobra.Command, args []string) {
	samepleConfig := config.DefaultAlertManagerConfig()
	fmt.Println(samepleConfig)
}

func init() {
	configCmd.AddCommand(generateTemplateCmd)
}
