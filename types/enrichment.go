package types

type Enrichment struct {
	StepName       string `yaml:"step_name"`
	EnrichmentName string `yaml:"enrichment_name"`
	EnrichmentArgs string `yaml:"enrichment_args"`
}

func GetDefaultEnrichment() Enrichment {
	return Enrichment{
		StepName:       "ENRICHMENT_STEP_1",
		EnrichmentName: "NOOP_ENRICHMENT",
		EnrichmentArgs: "ARG1,ARG2"}
}
