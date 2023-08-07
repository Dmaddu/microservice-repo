package main

import "github.com/gofiber/fiber/v2"

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Responding from Fiber Microservice!!")
	})

	// Start server on port 3000
	app.Listen(":3000")
}
