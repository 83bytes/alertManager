package action

import (
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
)

func NoopAction(a types.Action, resultMap map[string]interface{}) error {
	logr := logging.GetLogger()
	logr.Debug("noop action called")
	rs := fmt.Sprintf("noop action called \n")

	for k, v := range resultMap {
		if s, ok := v.(string); ok {
			rs += fmt.Sprintf("result of %s enrichment; is -> %s \n", k, s)
		}

	}

	logr.Debug(rs)
	return nil
}
