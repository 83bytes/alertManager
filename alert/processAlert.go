package alert

import (
	"alertmanager/config"
	"alertmanager/logging"
)

// Once I have an alertName
// I check for its pipeline in the AMconfig
// Once i get that
// I can the proper enrichmentFunc
// and the proper actionFunc
func ProcessAlert(a Alert) {
	logr := logging.GetLogger()
	an := a.GetAlertName()

	logr.Debug("alert processor", an)

	amc := config.GetAmConfig()
	p := amc.GetPipelineForAlert(an)
	if p != nil {
		logr.Infof("no alert-pipeline configured for %s", an)
	}

	logr.Debug("alert pipeline", p)

}
