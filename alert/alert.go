package alert

import (
	"alertmanager/types"
	"encoding/json"
	"fmt"
	"time"
)

type Alert struct {
	// fields we use to process stuff
	alertName   string             `json:"-"`
	enrichments []types.Enrichment `json:"-"`
	actions     []types.Action     `json:"-"`

	// fields we get from outside
	Annotations map[string]string `json:"annotations"`
	Labels      map[string]string `json:"labels"`
	StartsAt    time.Time         `json:"startsAt"`
	Status      string            `json:"status"`
}

type AlertGroup struct {
	Version           string            `json:"version"`
	GroupKey          string            `json:"groupKey"`
	TruncatedAlerts   int               `json:"truncatedAlerts"`
	Status            string            `json:"status"`
	Receiver          string            `json:"receiver"`
	GroupLabels       map[string]string `json:"groupLabels"`
	CommonLabels      map[string]string `json:"commonLabels"`
	CommonAnnotations map[string]string `json:"commonAnnotations"`
	ExternalURL       string            `json:"externalURL"`
	Alerts            []Alert           `json:"alerts"`
}

// getter and setter for internal fields
func (a *Alert) GetAlertName() string {
	return a.alertName
}

func (c AlertGroup) String() string {
	s, _ := json.Marshal(c)
	// we dont need to look at this error as we are marshalling a struct.
	// all error that can happen from loading random data into a struct are
	// handled at the ValidateAndLoad level
	return string(s)
}

func LoadAlertFromPayload(a *Alert) error {
	if an, ok := a.Labels["alertname"]; ok {
		a.alertName = an
		return nil
	}
	return fmt.Errorf("alertname not present in alert payload")
}
