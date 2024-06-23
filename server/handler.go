package server

import "github.com/gofiber/fiber/v2"

func alertWebhookHandler(c *fiber.Ctx) error {
	return c.SendString("pong")
}
