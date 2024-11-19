package routes

import (
	"gobby/src/logger"
	"time"

	"github.com/gofiber/fiber/v3"
)

func InitRoutes() {
	logger.Sugar.Info("initializing routes")
	app := fiber.New()

	// TODO assert parallelism behavior
	app.Get("/", func(c fiber.Ctx) error {
		// Send a string response to the client
		logger.Sugar.Info("Hello, wait 10s")
		time.Sleep(8 * time.Second)
		logger.Sugar.Info("Bye")
		return nil
	})

	// Start the server on port 3000
	logger.Sugar.Info("now listening on port 3000")
	app.Listen(":3000")
}
