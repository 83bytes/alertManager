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
