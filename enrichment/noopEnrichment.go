package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
)

func NoopEnrichment(e types.Enrichment) (interface{}, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprint("noop enrichment called with : ", e.EnrichmentArgs)
	logr.Info(rs)

	return rs, nil
}
