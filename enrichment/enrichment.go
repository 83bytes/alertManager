package enrichment

import "alertmanager/utils"

type Enrichment struct {
	EnrichmentName string `yaml:"enrichment_name"`
	EnrichmentArgs string `yaml:"enrichment_args"`
}

func GetDefaultEnrichment() Enrichment {
	return Enrichment{EnrichmentName: "NOOP_ENRICHMENT", EnrichmentArgs: "ARG1,ARG2"}
}

var enrichmentMap = make(utils.FunctionLut)

func GetEnrichmentMap() *utils.FunctionLut {
	return &enrichmentMap
}
