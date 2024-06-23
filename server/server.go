package server

import (
	"alertmanager/config"
	"fmt"
	"strconv"

	"github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	ServerPort     int
	MetricsPort    int
	ManagementPort int
	Config         *config.AlertManagerConfig
	Log            *logrus.Logger
}

func (s *Server) Start() error {
	fmt.Printf("Server start called %v", s)

	app := fiber.New()

	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Listen(":" + strconv.Itoa(s.ServerPort))
	return nil
}
