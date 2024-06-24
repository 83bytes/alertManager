package cmd

import (
	"alertmanager/config"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// validateCmd represents the validate command
var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "validate a config-file for errors",
	RunE:  validateCmdRunE,
}

func init() {
	configCmd.AddCommand(validateCmd)

	validateCmd.Flags().String("config-file", "./alert-manager-config.yml", "Path to config for validation")
}

func validateCmdRunE(cmd *cobra.Command, args []string) error {
	configFilePath, _ := cmd.Flags().GetString("config-file")

	b, err := os.ReadFile(configFilePath)
	if err != nil {
		return fmt.Errorf("cannot read config file; %s", err)
	}

	amConfig, err := config.ValidateAndLoad(b)
	if err != nil {
		return fmt.Errorf("validation failed: please refer to template; %s", err)
	}

	fmt.Println("config is correct; printing config")
	fmt.Println(amConfig)
	return nil
}
