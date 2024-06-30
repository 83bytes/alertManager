package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"os"
	"strings"
	"testing"
)

func TestNoopEnrichment(t *testing.T) {

	t.Run("basic noop enrichment", func(t *testing.T) {
		alert := types.Alert{AlertName: "TestAlert"}
		enrichment := types.Enrichment{EnrichmentName: "TestEnrichment", EnrichmentArgs: "enrichmentArg1,EnrichmentArg2"}

		log, err := logging.NewLogger("DEBUG")
		if err != nil {
			t.Errorf("error initializing logger")
		}
		nullFile, _ := os.OpenFile(os.DevNull, os.O_WRONLY, 0644)
		log.SetOutput(nullFile)
		result, _ := NoopEnrichment(alert, enrichment)

		if resultStr, ok := result.(string); ok {

			if !strings.Contains(resultStr, "noop enrichment called") {
				t.Errorf("expected result to contain 'noop enrichment called', got %s", resultStr)
			}

			if !strings.Contains(resultStr, "alert: TestAlert") {
				t.Errorf("expected result to contain 'alert: TestAlert', got %s", resultStr)
			}

			if !strings.Contains(resultStr, "enrichment: TestEnrichment") {
				t.Errorf("expected result to contain 'enrichment: TestEnrichment', got %s", resultStr)
			}

			if !strings.Contains(resultStr, "with args: enrichmentArg1,EnrichmentArg2") {
				t.Errorf("expected result to contain 'with args: enrichmentArg1,EnrichmentArg2', got %s", resultStr)
			}
		} else {
			t.Errorf("not returing a valid string got: %s, %s", result, resultStr)

		}

	})

}
