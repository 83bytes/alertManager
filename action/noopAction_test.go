package action

import (
	"alertmanager/logging"
	"alertmanager/types"
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestNoopAction(t *testing.T) {

	t.Run("basic noop action", func(t *testing.T) {
		alert := types.Alert{AlertName: "TestAlert"}
		action := types.Action{ActionName: "TestAction", ActionArgs: "actionArg1,actionArg3"}
		resultMap := map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		}

		log, err := logging.NewLogger("DEBUG")
		if err != nil {
			t.Errorf("error initializing logger")
		}
		var buf bytes.Buffer
		log.SetOutput(&buf)
		defer func() {
			log.SetOutput(os.Stderr)
		}()

		err = NoopAction(alert, action, resultMap)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}

		output := buf.String()
		if !strings.Contains(output, "noop action called") {
			t.Errorf("expected log to contain 'noop action called', got %s", output)
		}
		if !strings.Contains(output, "alert: TestAlert") {
			t.Errorf("expected log to contain 'alert: TestAlert', got %s", output)
		}
		if !strings.Contains(output, "action: TestAction") {
			t.Errorf("expected log to contain 'action: TestAction', got %s", output)
		}
		if !strings.Contains(output, "result of key1 enrichment(s):  value1") {
			t.Errorf("expected log to contain 'result of key1 enrichment(s):  value1', got %s", output)
		}
		if !strings.Contains(output, "result of key2 enrichment(s):  value2") {
			t.Errorf("expected log to contain 'result of key2 enrichment(s):  value2', got %s", output)
		}
	})

}
