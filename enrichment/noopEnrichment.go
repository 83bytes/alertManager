package enrichment

import (
	"alertmanager/logging"
	"fmt"
)

func NoopEnrichment(e Enrichment) (interface{}, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprint("noop enrichment called with : ", e.EnrichmentArgs)
	logr.Info(rs)

	return rs, nil
}
