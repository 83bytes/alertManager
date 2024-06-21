package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate a config-file for errors",
	Run:   validateCmdRun,
}

func init() {
	configCmd.AddCommand(validateCmd)

	validateCmd.Flags().String("config-file", "./alert-manager-config.yml", "Path to config for validation")
}

func validateCmdRun(cmd *cobra.Command, args []string) {
	// TODO:
	// Get file from command options
	// call validate on it
	fmt.Println("validate called")
}
