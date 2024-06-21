package config

import "fmt"

type AlertManagerConfig struct {
	AlertPipelines []AlertPipelineConfig
}

type AlertPipelineConfig struct{}

func (c AlertManagerConfig) String() string {
	return
}

func (c AlertManagerConfig) Validate() {
	fmt.Println("Validate Config")
}
