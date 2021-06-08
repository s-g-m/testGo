package server

import (
	"os"
	"test/server/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RunServer() *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	routes.Load(app)
	if os.Getenv("ENVIRONMENT") != "testing" {
		app.Listen(":3000")
	}
	return app
}
