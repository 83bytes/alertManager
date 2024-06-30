package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
)

func NoopEnrichment(alert types.Alert, e types.Enrichment) (interface{}, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprintf("noop enrichment called: \nalert: %s\nenrichment: %s\nwith args: %s", alert.AlertName, e.EnrichmentName, e.EnrichmentArgs)

	logr.Debug(rs)

	return rs, nil
}
