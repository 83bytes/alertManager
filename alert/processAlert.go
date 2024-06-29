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
// The body of the Action is passed to all the enrichment and action
// so that they have the complete context regarding what is going on
// The actions additionally will also contain the output of all the enrichments
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
	var err error
	resMap := make(map[string]interface{}, len(p.Enrichments))

	// process enrichments
	for _, v := range (*p).Enrichments {
		logr.Info("processing enrichment : ", v.EnrichmentName)

		if f, ok := (*enrichmentMap)[v.EnrichmentName]; ok {
			resMap[v.EnrichmentName], err = f(v)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	// process actions
	for _, v := range (*p).Actions {
		logr.Info("processing action : ", v.ActionName)

		if f, ok := (*actionMap)[v.ActionName]; ok {
			err := f(v, resMap)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}
