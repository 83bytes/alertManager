package enrichment

import (
	"alertmanager/logging"
	"alertmanager/types"
	"strings"
	"testing"
)

func TestUpperCaseEnrichment(t *testing.T) {

	t.Run("upper case enrichment", func(t *testing.T) {
		alert := types.Alert{AlertName: "TestAlert"}
		enrichment := types.Enrichment{EnrichmentName: "UPPER_CASE", EnrichmentArgs: "enrichmentArg1,enrichmentArg2"}

		_, err := logging.NewLogger("DEBUG")
		if err != nil {
			t.Errorf("error initializing logger")
		}

		result, _ := UpperCaseEnrichment(alert, enrichment)

		if resultStr, ok := result.(string); ok {
			if !strings.Contains(resultStr, "ENRICHMENTARG1,ENRICHMENTARG2") {
				t.Errorf("expected result to contain 'ENRICHMENTARG1,ENRICHMENTARG2', got %s", resultStr)
			}
		} else {
			t.Errorf("not returing a valid string got: %s, %s", result, resultStr)

		}

	})

}
