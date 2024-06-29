package types

import (
	"encoding/json"
	"time"
)

type Alert struct {
	// fields we use to process stuff
	AlertName   string       `json:"-"`
	Enrichments []Enrichment `json:"-"`
	Actions     []Action     `json:"-"`

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
	return a.AlertName
}

func (c AlertGroup) String() string {
	s, _ := json.Marshal(c)
	// we dont need to look at this error as we are marshalling a struct.
	// all error that can happen from loading random data into a struct are
	// handled at the ValidateAndLoad level
	return string(s)
}
