package middleware

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

func AuthenticateUser(c *fiber.Ctx) error {
	// print current date and time to console
	fmt.Println("Date:", time.Now(), " User Authenticated")

	// passing the request to the next endpoint
	return c.Next()
}
