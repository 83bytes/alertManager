package cmd

import (
	"alertmanager/action"
	"alertmanager/config"
	"alertmanager/enrichment"
	"alertmanager/logging"
	"alertmanager/server"
	"fmt"
	"os"

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
	cFile, _ := cmd.Flags().GetString("config-file")

	log, err := logging.NewLogger(ll)
	if err != nil {
		return err
	}

	// Initialize the Enrichments
	log.Info("Adding NooP Enrichment")
	enr := enrichment.GetEnrichmentMap()
	enr.Add("NOOP_ENRICHMENT", enrichment.NoopEnrichment)

	// Initialize the Actions
	log.Info("Adding NooP Action")
	actMap := action.GetActionMap()
	actMap.Add("NOOP_ACTION", action.NoopAction)

	b, err := os.ReadFile(cFile)
	if err != nil {
		return fmt.Errorf("cannot read config file; %s", err)
	}
	amConfig, err := config.ValidateAndLoad(b)
	if err != nil {
		return fmt.Errorf("validation failed: please refer to template; %s", err)
	}
	fmt.Println("printing config")
	fmt.Println(amConfig)

	var s = server.Server{
		ServerPort:     sPort,
		MetricsPort:    mPort,
		ManagementPort: mgmtPort,
		Config:         amConfig,
		Log:            log,
	}

	s.Start()
	return nil
}

func init() {
	rootCmd.AddCommand(serverCmd)

	serverCmd.Flags().Int("server-port", 8081, "Port to listen on")
	// todo set up management api and metrics endpoint
	// serverCmd.Flags().Int("metric-port", 8082, "metrics port to listen on")
	// serverCmd.Flags().Int("management-port", 8083, "management port to listen on")
	serverCmd.Flags().String("config-file", "./alert-manager-config.yml", "Path to alert config")
	serverCmd.Flags().String("log-level", logging.DEFAULT_LOG_LEVEL, "log-level for alertmanager; options INFO|DEBUG|ERROR")
}
