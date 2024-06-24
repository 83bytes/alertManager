package server

import (
	"alertmanager/config"
	"alertmanager/logging"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Server struct {
	ServerPort     int
	MetricsPort    int
	ManagementPort int
	Config         *config.AlertManagerConfig
	Log            *logging.Logger
}

func (s *Server) Start() error {
	app := fiber.New()

	// handle gradeful app shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt) // SIGINT

	var serverShutDownWG sync.WaitGroup // wait for the server to shutdown cleanly

	// shutdown server upon SIGINT and signal the waitgroup
	go func() {
		<-c

		s.Log.Info("shutting down server cleanly; please wait :)")
		serverShutDownWG.Add(1)
		defer serverShutDownWG.Done()

		_ = app.ShutdownWithTimeout(time.Second * 30)
	}()

	// request logger middlewear
	// should be okay as long the server is low-volume
	// maybe we can configure to log only bad / failed reqs
	app.Use(logger.New())

	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	app.Post("/webhook/", alertWebhookHandler)

	logr := logging.GetLogger()
	logr.Info("Starting Server ...")
	if err := app.Listen(":" + strconv.Itoa(s.ServerPort)); err != nil {
		s.Log.Error("server: unable to start server", err)
	}

	serverShutDownWG.Wait()
	return nil
}
