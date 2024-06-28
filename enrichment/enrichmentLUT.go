package enrichment

type EnrichmentLut map[string]func(Enrichment) (interface{}, error)

func (flut EnrichmentLut) Add(fname string, f func(Enrichment) (interface{}, error)) {
	flut[fname] = f
}
