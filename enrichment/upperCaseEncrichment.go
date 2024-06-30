package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
	"strings"
)

func UpperCaseEnrichment(alert types.Alert, e types.Enrichment) (interface{}, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprintf("upper-case enrichment called: \nalert: %s\nenrichment: %s\nwith args: %s", alert.AlertName, e.EnrichmentName, e.EnrichmentArgs)

	logr.Debug(rs)

	return strings.ToUpper(e.EnrichmentArgs), nil
}
