package main

import (
	"github.com/Dmaddu/microservice-repo/internal/app/middleware"
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	//app.Use(middleware.AuthenticateUser)
	app.Get("/", middleware.AuthenticateUser, func(c *fiber.Ctx) error {
		return c.SendString("Responding from Fiber Microservice!")
	})

	// Start server on port 3000
	app.Listen(":3000")
}
