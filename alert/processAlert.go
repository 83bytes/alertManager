package alert

import (
	"alertmanager/action"
	"alertmanager/config"
	"alertmanager/enrichment"
	"alertmanager/logging"
	"fmt"
)

// Processes an entire alert-pipeline end-to-end
// If an alert-pipeline is configured for given alert
// this function executes the enrichments one-by-one
// then it executes the actions one-by-one
// The body of the Action is passed to all the enrichment and action
// so that they have the complete context regarding what is going on
// The actions additionally will also contain the output of all the enrichments
func ProcessAlert(a Alert) {
	logr := logging.GetLogger()
	an := a.GetAlertName()

	logr.Debug("alert processor", an)

	amc := config.GetAmConfig()
	p := amc.GetPipelineForAlert(an)
	if p == nil {
		logr.Infof("no alert-pipeline configured for %s", an)
		return
	}
	logr.Debug("alert pipeline name", (*p))

	enrichmentMap := enrichment.GetEnrichmentMap()
	actionMap := action.GetActionMap()

	// these are used to capture the result of the enrichments
	enrichmentRes := ""
	var err error

	for _, v := range (*p).Enrichments {
		logr.Info("processing enrichment : ", v.EnrichmentName)

		if f, ok := (*enrichmentMap)[v.EnrichmentName]; ok {
			enrichmentRes, err = f(v.EnrichmentArgs)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("********** ", enrichmentRes, "*****")
		}
	}

	fmt.Print("ENRICHMENT RESP ************* ", enrichmentRes)

	for _, v := range (*p).Actions {
		logr.Info("processing action : ", v.ActionName)

		if f, ok := (*actionMap)[v.ActionName]; ok {
			x, err := f(v.ActionArgs + enrichmentRes)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("################## ", x)
		}
	}

}
