package utils

import "github.com/gofiber/fiber/v2"

func HealthCheck(app *fiber.App) {
	// Health Check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
	   // Check your dependencies (database, cache, etc.) here.
	   // If everything is okay, return a success. Otherwise, return an error.
	   
	   // For this simple example, we're just returning a success without any checks.
	   return c.Status(fiber.StatusOK).JSON(fiber.Map{
		   "status": "success",
		   "message": "All systems operational",
	   })
   })
}