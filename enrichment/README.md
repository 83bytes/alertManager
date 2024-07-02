# Enrichment

Enrichments are a way to supercharge your alerts by gather relevant details about your system to better understand the alert-context.

These can be anything, from grabbing a quick chart to see overall CPU-utilization to getting a heap-dump from a crashing pod.

## How to build enrichments

### Function Signature and Data-types

We build enrichments by writing enrichment functions in golang which satisfy the type

```
func (types.Alert, types.Enrichment) (interface{}, error)
```

This is the function that the tam-runtime will call once we register this enrichment function

The enrichment type looks like this

```
type Enrichment struct {
	StepName       string `yaml:"step_name"`
	EnrichmentName string `yaml:"enrichment_name"`
	EnrichmentArgs string `yaml:"enrichment_args"`
}
```

### Program context

Once we are inside the function-context, we are free to do anything we want. <br>
We have the entire alert that was used to trigger this enrichment.

The environment takes care of storing the output and passing it to actions later in the pipeline.

**NOTE:** Enrichment Functions do not share context with each other. Thus you **CANNOT** use the output of one enrichment in another enrichment.

### Registering the enrichment-function

Once we have defined out function, we have to register it with the tam-runtime. This is basically updating a in-memory map which stores all the enrichments available and a string identifier that is used to identify this function.

The function in [enrichment.go](./enrichment.go) looks like this

```
func LoadEnrichments() {
	enr := GetEnrichmentMap()
	enr.Add("NOOP_ENRICHMENT", NoopEnrichment)
	enr.Add("UPPER_CASE", UpperCaseEnrichment)
	enr.Add("GetPromQL", GetPromQLEnrichment)
}
```

we will add new entry in this function which like this

```
enr.Add("EnrichmentIdentifier", EnrichmentFunctionName)
```

Here,
`EnrichmentIdentifier` is the string value that will be used to refer to this enrichment in the tam-config; `EnrichmentFunctionName` is the name of the function that we defined
