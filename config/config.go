package config

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type AlertManagerConfig struct {
	AlertPipelines []AlertPipelineConfig `yaml:"alert_pipelines"`
}

type AlertPipelineConfig struct {
	AlertName   string       `yaml:"alert_name"`
	Enrichments []Enrichment `yaml:"enrichments"`
	Actions     []Action     `yaml:"actions"`
}

type Enrichment struct {
	EnrichmentName string `yaml:"enrichment_name"`
	EnrichmentArgs string `yaml:"enrichment_args"`
}

type Action struct {
	ActionName string `yaml:"action_name"`
	ActionArgs string `yaml:"action_args"`
}

func defaultAction() Action {
	return Action{ActionName: "NOOP_ACTION", ActionArgs: "ARG1,ARG2"}
}

func defaultEnrichment() Enrichment {
	return Enrichment{EnrichmentName: "NOOP_ENRICHMENT", EnrichmentArgs: "ARG1,ARG2"}
}

func defaultAlertPipelineConfig() AlertPipelineConfig {
	return AlertPipelineConfig{
		AlertName:   "NOOP_ALERT",
		Enrichments: []Enrichment{defaultEnrichment()},
		Actions:     []Action{defaultAction()},
	}
}

func DefaultAlertManagerConfig() AlertManagerConfig {
	return AlertManagerConfig{
		AlertPipelines: []AlertPipelineConfig{defaultAlertPipelineConfig()},
	}
}

func (c AlertManagerConfig) String() string {
	s, err := yaml.Marshal(c)
	if err != nil {
		fmt.Println("eror", err)
		// TODO: handle properly
	}
	return string(s)
}

func (c AlertManagerConfig) Validate() error {
	fmt.Println("Validate Config")
	fmt.Println(c)
	return nil
	// TODO: add validation code
	// Returns Count of rules present, error
}
