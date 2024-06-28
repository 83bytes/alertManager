package enrichment

import (
	"alertmanager/logging"
	"fmt"
)

func NoopEnrichment(args string) (string, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprint("noop enrichment called with : ", args)
	logr.Info(rs)

	return rs, nil
}
