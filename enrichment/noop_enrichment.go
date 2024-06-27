package enrichment

import (
	"alertmanager/logging"
)

func Noop_Enrichment(args string) (string, error) {
	logr := logging.GetLogger()

	logr.Info("Noop enrichment Called")

	return "Noop called", nil
}
