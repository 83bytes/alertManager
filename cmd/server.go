package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the AlertManager Webhook Server",
	RunE:  serverCommandRunE,
}

func serverCommandRunE(cmd *cobra.Command, args []string) error {
	fmt.Println("Server is being called")
	ll, _ := cmd.Flags().GetString("log-level")
	fmt.Println(ll)
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serverCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serverCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
