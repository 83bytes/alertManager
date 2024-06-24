package alert

import "alertmanager/logging"

func ProcessAlert(a Alert) {
	logr := logging.GetLogger()
	an := a.GetAlertName()

	logr.Debug("alert processor", an)

}
