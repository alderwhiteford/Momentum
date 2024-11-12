package milestoneService

import "github.com/gofiber/fiber/v2"

func (c *MilestoneServiceImpl) InitializeRoutes(app fiber.Router, adminMiddleware fiber.Handler) {
	milestone := app.Group("/milestone")

	// ADMIN ROUTES:
	milestone.Get("/")
}