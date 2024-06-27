package alert

import (
	"alertmanager/config"
	"alertmanager/enrichment"
	"alertmanager/logging"
	"fmt"
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
	if p == nil {
		logr.Infof("no alert-pipeline configured for %s", an)
		return
	}
	logr.Debug("alert pipeline name", (*p))

	enrichmentMap := enrichment.GetEnrichmentMap()

	for _, v := range (*p).Enrichments {
		logr.Info("processing enrichment : ", v.EnrichmentName)

		if f, ok := (*enrichmentMap)[v.EnrichmentName]; ok {
			x, err := f(v.EnrichmentArgs)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Print("********** ", x)
		}

	}

}
