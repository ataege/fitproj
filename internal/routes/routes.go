package routes

import (
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes configures all the routes for the application
func SetupRoutes(app *fiber.App) {
	public := app.Group("/api/v1")
	setupPublicRoutes(public)
}

func setupPublicRoutes(router fiber.Router) {
	router.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status":  "success",
			"message": "Service is healthy",
		})
	})
}
