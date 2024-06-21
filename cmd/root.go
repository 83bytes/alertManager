package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var version = "0.0.1"
var DEFAULT_LOG_LEVEL = "INFO"

var rootCmd = &cobra.Command{
	Use:     "alertmanager",
	Version: version,
	Short:   "An alertmanager for managing alerts",
	Long: `This alertmanager tool can be used to start the alertmanager server or to validate the config.
Additionally you can also generate a bare config-template`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().String("log-level", DEFAULT_LOG_LEVEL, "log-level for alertmanager; options INFO|DEBUG|ERROR")
}
