package goalService

import (
	"github.com/gofiber/fiber/v2"
)

// This method will consume the users/{user_id} grouping:
func (c *GoalServiceImpl) InitializeRoutes(app fiber.Router) {
	goal := app.Group("/goal");
	
	goal.Post("/", c.CreateGoal)
}