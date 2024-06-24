package config

import (
	"alertmanager/action"
	"alertmanager/enrichment"
	"fmt"

	"gopkg.in/yaml.v3"
)

type AlertManagerConfig struct {
	AlertPipelines []AlertPipelineConfig `yaml:"alert_pipelines"`
}

type AlertPipelineConfig struct {
	AlertName   string                  `yaml:"alert_name"`
	Enrichments []enrichment.Enrichment `yaml:"enrichments"`
	Actions     []action.Action         `yaml:"actions"`
}

func defaultAlertPipelineConfig() AlertPipelineConfig {
	return AlertPipelineConfig{
		AlertName:   "NOOP_ALERT",
		Enrichments: []enrichment.Enrichment{enrichment.GetDefaultEnrichment()},
		Actions:     []action.Action{action.GetDefaultAction()},
	}
}

func DefaultAlertManagerConfig() AlertManagerConfig {
	return AlertManagerConfig{
		AlertPipelines: []AlertPipelineConfig{defaultAlertPipelineConfig()},
	}
}

func (c AlertManagerConfig) String() string {
	s, _ := yaml.Marshal(c)
	// we dont need to look at this error as we are marshalling a struct.
	// all error that can happen from loading random data into a struct are
	// handled at the ValidateAndLoad level
	return string(s)
}

func ValidateAndLoad(b []byte) (AlertManagerConfig, error) {
	amConfig := AlertManagerConfig{}

	// todo
	// try to use a strict unmarshalling like in json
	err := yaml.Unmarshal(b, &amConfig)
	if err != nil {
		return AlertManagerConfig{},
			fmt.Errorf("unable to load config, please check format; %s", err)
	}

	if len(b) > 0 && amConfig.AlertPipelines == nil {
		return AlertManagerConfig{},
			fmt.Errorf("unable to load config, please check format")
	}
	// todo:
	// do better validation
	// right now it accepts a stray key in the list of alert_pipelines
	// and ingects an empty alert-config
	// Filter out the empty entr for now.
	// maybe check if json-schema etc can help here
	return amConfig, nil
}
