package main

import (
	// "github.com/gofiber/fiber/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()

	// Use the logger middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/", func(c *fiber.Ctx) error {
		// panic("err")
		return c.Status(200).JSON(fiber.Map{
			"message": "pong",
		})
	})

	app.Listen(":8080")
}
