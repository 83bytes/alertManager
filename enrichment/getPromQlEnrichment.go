package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Result represents an individual result in the data
type Result struct {
	Metric map[string]string `json:"metric"`
	Value  []interface{}     `json:"value"`
}

// Data represents the data section of the response
type Data struct {
	ResultType string   `json:"resultType"`
	Result     []Result `json:"result"`
}

// Response represents the entire response structure
type PromResponse struct {
	Status string `json:"status"`
	Data   Data   `json:"data"`
}

func GetPromQLEnrichment(alert types.Alert, e types.Enrichment) (interface{}, error) {
	logr := logging.GetLogger()

	logr.Debug("getting promql data form endpoint")

	resp, err := http.Get(e.EnrichmentArgs)
	if err != nil {
		return nil, fmt.Errorf("error reaching server at: %s", e.EnrichmentArgs)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-200 response status: %d", resp.StatusCode)
	}
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}
	var pres PromResponse
	err = json.Unmarshal([]byte(bodyBytes), &pres)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	// check if result is empty
	if len(pres.Data.Result) <= 0 {
		return "", nil
	}
	return fmt.Sprint(pres.Data.Result[0].Value[1]), nil
}
