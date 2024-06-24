package server

import (
	"alertmanager/alert"
	"alertmanager/logging"
	"alertmanager/utils"

	"github.com/gofiber/fiber/v2"
)

func alertWebhookHandler(c *fiber.Ctx) error {
	logr := logging.GetLogger()
	logr.Debug("in webhook handler")
	ag := new(alert.AlertGroup)

	b := c.BodyRaw()

	err := utils.StrictUnmarshal(b, ag)
	if err != nil {
		logr.Debugf("Body being processed %s", string(b))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// if any alert does not have alertname set as a label,
	// we reject it as unprocessable
	// this is because we rely on this label to select and match
	// enrichments and actions
	alerts := ag.Alerts
	for i := 0; i < len(alerts); i++ {
		logr.Debugf("Processing Alert %d -> %s", i, alerts[i])
		a := alerts[i]
		err := alert.LoadAlertFromPayload(&a)
		if err != nil {
			c.SendStatus(fiber.StatusBadRequest)
		}
		logr.Debugf("%s", a.GetAlertName())
	}
	return c.SendStatus(fiber.StatusOK)
}
