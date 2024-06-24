package config

import (
	"reflect"
	"testing"
)

var (
	goodSingleAlertConfig = `
alert_pipelines:
  - alert_name: NOOP_ALERT
    enrichments:
    - enrichment_name: NOOP_ENRICHMENT
      enrichment_args: ARG1,ARG2
    actions:
    - action_name: NOOP_ACTION
      action_args: ARG1,ARG2
`

	goodDoubleAlertConfig = `
alert_pipelines:
  - alert_name: NOOP_ALERT
    enrichments:
    - enrichment_name: NOOP_ENRICHMENT
      enrichment_args: ARG1,ARG2
    actions:
    - action_name: NOOP_ACTION
      action_args: ARG1,ARG2
  - alert_name: NOOP_ALERT
    enrichments:
    - enrichment_name: NOOP_ENRICHMENT
      enrichment_args: ARG1,ARG2
    actions:
    - action_name: NOOP_ACTION
      action_args: ARG1,ARG2
`
)

func TestValidateAndLoad(t *testing.T) {
	type args struct {
		b []byte
	}

	tests := []struct {
		name    string
		args    args
		want    AlertManagerConfig
		wantErr bool
	}{
		{
			name: "Good default alert config",
			args: args{
				b: []byte(goodSingleAlertConfig),
			},
			want:    DefaultAlertManagerConfig(),
			wantErr: false,
		},
		{
			name: "Good double alert config",
			args: args{
				b: []byte(goodDoubleAlertConfig),
			},
			want: AlertManagerConfig{
				AlertPipelines: []AlertPipelineConfig{
					defaultAlertPipelineConfig(),
					defaultAlertPipelineConfig(),
				},
			},
			wantErr: false,
		},
		{
			name: "Random Yaml",
			args: args{
				b: []byte(`randomKey: randonValue
		randomKey2: randomValue2`),
			},
			want:    AlertManagerConfig{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ValidateAndLoad(tt.args.b)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateAndLoad() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// dereference the pointer for got to make deepequal happy
			if !reflect.DeepEqual(*got, tt.want) {
				t.Errorf("ValidateAndLoad() = \n%v, want \n%v", got, tt.want)
			}
		})
	}
}
