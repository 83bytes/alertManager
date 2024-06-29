package config

import (
	"alertmanager/types"
	"fmt"

	"gopkg.in/yaml.v3"
)

// we create a global instance of AlertManagerConfig

var AmConfig = new(types.AlertManagerConfig)

func GetAmConfig() *types.AlertManagerConfig {
	return AmConfig
}

func ValidateAndLoad(b []byte) (*types.AlertManagerConfig, error) {
	amConfig := GetAmConfig()

	// todo protect this by a Mutex
	// a write mutex is enough
	// todo
	// try to use a strict unmarshalling like in json
	err := yaml.Unmarshal(b, &amConfig)
	if err != nil {
		return &types.AlertManagerConfig{},
			fmt.Errorf("unable to load config, please check format; %s", err)
	}

	if len(b) > 0 && amConfig.AlertPipelines == nil {
		return &types.AlertManagerConfig{},
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
