package goalService

import (
	"github.com/gofiber/fiber/v2"
)

// This method will consume the users/{user_id} grouping:
func (c *GoalServiceImpl) InitializeRoutes(app fiber.Router, adminMiddleware fiber.Handler) {
	goal := app.Group("/goal")
	goalById := goal.Group("/:goalId")

	// ADMIN ROUTES:
	goal.Get("/", adminMiddleware, c.GetAllGoals)

	// ENTITY SPECIFIC:
	goalById.Get("/", c.GetGoal)
	goalById.Delete("/", c.DeleteGoal)
	goalById.Patch("/", c.UpdateGoal)
}