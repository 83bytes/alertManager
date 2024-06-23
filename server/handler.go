package server

import (
	"github.com/gofiber/fiber/v2"
)

func alertWebhookHandler(c *fiber.Ctx) error {
	c.Context().Logger()
	b := c.BodyRaw()
	return c.Status(fiber.StatusBadRequest).Send(b)
}
