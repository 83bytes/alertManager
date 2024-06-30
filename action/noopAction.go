package action

import (
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
)

func NoopAction(alert types.Alert, action types.Action, resultMap map[string]interface{}) error {
	logr := logging.GetLogger()

	rs := fmt.Sprintf("noop action called \nalert: %s \naction: %s \n", alert.AlertName, action.ActionName)

	for k, v := range resultMap {
		if s, ok := v.(string); ok {
			rs += fmt.Sprintf("result of %s enrichment(s):  %s \n", k, s)
		}
	}

	logr.Debug(rs)
	return nil
}
