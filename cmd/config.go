package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
// there are 2 subcommands
// validate: used to validate a given config-gile
// generate-template: used to generate a blank template that is filled with defaults
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "Use this command to validate an existing config-file or to generate a sample template",
}

func init() {
	rootCmd.AddCommand(configCmd)
}
