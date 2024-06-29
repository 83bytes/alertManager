package enrichment

import "alertmanager/types"

var enrichmentMap = make(EnrichmentLut)

func GetEnrichmentMap() *EnrichmentLut {
	return &enrichmentMap
}

type EnrichmentLut map[string]func(types.Enrichment) (interface{}, error)

func (flut EnrichmentLut) Add(fname string, f func(types.Enrichment) (interface{}, error)) {
	flut[fname] = f
}
