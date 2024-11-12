package userService

import (
	"github.com/gofiber/fiber/v2"
)

func (c *UserServiceImpl) InitializeRoutes(app *fiber.App, userMiddleware fiber.Handler, adminMiddleware fiber.Handler) {
	user := app.Group("/user");

	// ADMIN ROUTES:
	user.Get("/", adminMiddleware, c.GetAllUsers)

	// ENTITY SPECIFIC
	userById := user.Group("/:userId")
	userById.Use(userMiddleware)

	userById.Get("/", c.GetUser)
	userById.Patch("/", c.UpdateUser)
	userById.Delete("/", c.DeleteUser)

	// ENTITY RELATIONSHIPS
	userGoals := userById.Group("/:goalId")

	userGoals.Get("/", c.GetUserGoals)
	userGoals.Post("/", c.CreateUserGoal)
}