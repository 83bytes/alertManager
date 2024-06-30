package enrichment

import "alertmanager/types"

var enrichmentMap = make(EnrichmentLut)

func GetEnrichmentMap() *EnrichmentLut {
	return &enrichmentMap
}

type EnrichmentFunc func(types.Alert, types.Enrichment) (interface{}, error)
type EnrichmentLut map[string]EnrichmentFunc

func (flut EnrichmentLut) Add(fname string, f EnrichmentFunc) {
	flut[fname] = f
}

// Use this function to load all the defined enrichments in memory
// is not goroutine safe
// todo: protect this with a mutex/sync.Once
func LoadEnrichments() {
	enr := GetEnrichmentMap()
	enr.Add("NOOP_ENRICHMENT", NoopEnrichment)
	enr.Add("UPPER_CASE", UpperCaseEnrichment)

}
