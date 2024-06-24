package alert

import (
	"alertmanager/config"
	"alertmanager/logging"
)

func ProcessAlert(a Alert) {
	logr := logging.GetLogger()
	an := a.GetAlertName()

	logr.Debug("alert processor", an)

	logr.Debug("Here is the alert config", config.GetAmConfig())

}
