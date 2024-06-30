package types

import (
	"alertmanager/logging"

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

func DefaultAlertPipelineConfig() AlertPipelineConfig {
	return AlertPipelineConfig{
		AlertName:   "NOOP_ALERT",
		Enrichments: []Enrichment{GetDefaultEnrichment()},
		Actions:     []Action{GetDefaultAction()},
	}
}

func DefaultAlertManagerConfig() AlertManagerConfig {
	return AlertManagerConfig{
		AlertPipelines: []AlertPipelineConfig{DefaultAlertPipelineConfig()},
	}
}
func (c AlertManagerConfig) String() string {
	s, _ := yaml.Marshal(c)
	// we dont need to look at this error as we are marshalling a struct.
	// all error that can happen from loading random data into a struct are
	// handled at the ValidateAndLoad level
	return string(s)
}

func (am *AlertManagerConfig) GetPipelineForAlert(name string) *AlertPipelineConfig {
	logr := logging.GetLogger()
	for _, pipes := range am.AlertPipelines {
		if pipes.AlertName == name {
			logr.Debug("pipeline found for alert : ", name)
			return &pipes
		}
	}
	logr.Debug("no pipelines found for alert : ", name)
	return nil
}
