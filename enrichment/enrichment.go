package enrichment

type Enrichment struct {
	EnrichmentName string `yaml:"enrichment_name"`
	EnrichmentArgs string `yaml:"enrichment_args"`
}

func GetDefaultEnrichment() Enrichment {
	return Enrichment{EnrichmentName: "NOOP_ENRICHMENT", EnrichmentArgs: "ARG1,ARG2"}
}
