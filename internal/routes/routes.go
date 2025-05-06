package routes

import (
	"github.com/ataege/fitproj/internal/handlers"
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

	// File routes
	router.Post("/upload", handlers.UploadFile)
	router.Post("/upload/multiple", handlers.UploadMultipleFiles)
	router.Get("/files", handlers.ListFiles)
	router.Get("/files/:filename", handlers.DownloadFile)
}
