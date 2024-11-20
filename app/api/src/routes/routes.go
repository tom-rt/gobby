package routes

import (
	"gobby/src/logger"

	"github.com/gofiber/fiber/v3"
)

func InitRoutes() {
	logger.Sugar.Info("initializing routes")
	app := fiber.New()

	app.Get("/ping", func(c fiber.Ctx) error {
		logger.Sugar.Info("ping request received")
		return c.SendString("pong")
	})

	// Start the server on port 3000
	logger.Sugar.Info("now listening on port 3000")
	app.Listen(":3000")
}
