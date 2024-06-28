package action

import (
	"alertmanager/logging"
	"fmt"
)

func NoopAction(args string) (string, error) {
	logr := logging.GetLogger()

	rs := fmt.Sprint("Noop Action Called; with output of enrichment ", args)
	logr.Info(rs)
	return rs, nil
}
