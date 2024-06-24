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
		logr.Debugf("body being processed %s", string(b))
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// if any alert does not have alertname set as a label,
	// we reject the entire payload as unprocessable
	// this is because we rely on this label to select and match
	// enrichments and actions
	alerts := ag.Alerts
	nalerts := len(alerts)
	for i := 0; i < nalerts; i++ {
		logr.Debugf(`loading alert %d -> %v`, i, alerts[i])
		a := alerts[i]
		err := alert.LoadAlertFromPayload(&a)
		if err != nil {
			return c.SendStatus(fiber.StatusBadRequest)
		}
		logr.Debugf("%s", a.GetAlertName())
	}

	// now that we know that all the alerts in the payload are processable
	// we can start to process them
	// todo
	// figure out the proper http response for a partial acceptance case
	// where we tell the server which alert in the payload is bad ??
	// todo
	// do we want to use go-routines and channels here ?
	// ideally would depend on the volume here
	for i := 0; i < nalerts; i++ {
		logr.Debugf(`processing alert %d -> %v`, i, alerts[i])
		alert.ProcessAlert(alerts[i])
	}

	return c.SendStatus(fiber.StatusOK)
}
