package action

import (
	"alertmanager/logging"
	"alertmanager/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type SlackMessage struct {
	Text string `json:"text"`
}

func SendToSlackAction(alert types.Alert, action types.Action, resultMap map[string]interface{}) error {
	logr := logging.GetLogger()
	logr.Info("sending data to slack")

	rs := fmt.Sprintf("alert: %s \naction: %s \n", alert.AlertName, action.ActionName)

	for k, v := range resultMap {
		if s, ok := v.(string); ok {
			rs += fmt.Sprintf("result of %s enrichment(s):  %s \n", k, s)
		}
	}

	logr.Debug("message being sent to slack", rs)

	slackMessage := SlackMessage{
		Text: rs,
	}

	jsonPayload, err := json.Marshal(slackMessage)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}

	req, err := http.NewRequest("POST", action.ActionArgs, bytes.NewBuffer(jsonPayload))
	if err != nil {
		return fmt.Errorf("failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send HTTP request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := io.ReadAll(resp.Body)
		bodyString := string(bodyBytes)
		return fmt.Errorf("received non-200 response status: %d, body: %s", resp.StatusCode, bodyString)
	}

	return nil
}
