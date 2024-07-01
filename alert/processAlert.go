package alert

import (
	"alertmanager/action"
	"alertmanager/config"
	"alertmanager/enrichment"
	"alertmanager/logging"
	"alertmanager/types"
	"fmt"
)

func LoadAlertFromPayload(a *types.Alert) error {
	if an, ok := a.Labels["alertname"]; ok {
		a.AlertName = an
		return nil
	}
	return fmt.Errorf("alertname not present in alert payload")
}

// Processes an entire alert-pipeline end-to-end
// If an alert-pipeline is configured for given alert
// this function executes the enrichments one-by-one
// then it executes the actions one-by-one
// The body of the Alert is passed to all the enrichment and action
// so that they have the complete context regarding what is going on
// The actions additionally will also contain the output of all the enrichments

// todo:
// actions and enrichments are not able to pick "specific information" from the Alert.
// want to add the ability to pick x.y.z field from the alert json.
// this will further power things like "lookup ip from alert and grab x,y,z metrics from grafana"

// todo:
// process each enrichment concurrently using goroutines and channels etc

func ProcessAlert(a types.Alert) {
	logr := logging.GetLogger()
	an := a.GetAlertName()

	logr.Info("alert being processed : ", an)

	amc := config.GetAmConfig()
	p := amc.GetPipelineForAlert(an)
	if p == nil {
		logr.Infof("no alert-pipeline configured for %s", an)
		return
	}
	logr.Debug("alert pipeline found for ", (*p))

	enrichmentMap := enrichment.GetEnrichmentMap()
	actionMap := action.GetActionMap()

	// these are used to capture the result of the enrichments
	resMap := make(map[string]interface{}, len(p.Enrichments))

	// process enrichments
	for _, v := range (*p).Enrichments {
		logr.Info("processing enrichment : ", v.EnrichmentName)

		if f, ok := (*enrichmentMap)[v.EnrichmentName]; ok {
			tt, err := f(a, v)
			if err != nil {
				logr.Error("error from processing function")
			}
			resMap[v.StepName] = tt
		} else {
			logr.Info("no enrichment found with name: ", v.EnrichmentName)
		}
	}

	// process actions
	for _, v := range (*p).Actions {
		logr.Info("processing action : ", v.ActionName)

		if f, ok := (*actionMap)[v.ActionName]; ok {
			err := f(a, v, resMap)
			if err != nil {
				logr.Error("error processing action:", err)
			}
		} else {
			logr.Info("no action found with name: ", v.ActionName)
		}
	}
}
