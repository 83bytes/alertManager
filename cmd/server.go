package cmd

import (
	"alertmanager/config"
	"alertmanager/server"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var serverCmd = &cobra.Command{
	Use:   "server",
	Short: "Start the AlertManager Webhook Server",
	RunE:  serverCommandRunE,
}

func serverCommandRunE(cmd *cobra.Command, args []string) error {
	ll, _ := cmd.Flags().GetString("log-level")
	sPort, _ := cmd.Flags().GetInt("server-port")
	mPort, _ := cmd.Flags().GetInt("metric-port")
	mgmtPort, _ := cmd.Flags().GetInt("meanagement-port")
	// _, _ := cmd.Flags().GetString("config-file")

	log := logrus.New()
	err := setLogLevelE(log, ll)
	if err != nil {
		return err
	}

	var s = server.Server{
		ServerPort:     sPort,
		MetricsPort:    mPort,
		ManagementPort: mgmtPort,
		Config:         &config.AlertManagerConfig{},
		Log:            log,
	}

	s.Start()
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Int("server-port", 8081, "Port to listen on")
	serverCmd.Flags().Int("metric-port", 8082, "metrics port to listen on")
	serverCmd.Flags().Int("management-port", 8083, "management port to listen on")
	serverCmd.Flags().String("config-file", "./alert-manager-config.yml", "Path to alert config")
	serverCmd.Flags().String("log-level", DEFAULT_LOG_LEVEL, "log-level for alertmanager; options INFO|DEBUG|ERROR")
}
