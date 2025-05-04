package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func GetHealth(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "healthy",
	})
}

// GetHello returns a simple hello message
func GetHello(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Hello, World!",
	})
}
