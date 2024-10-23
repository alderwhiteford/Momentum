package userService

import (
	"github.com/gofiber/fiber/v2"
)

func (c *UserServiceImpl) InitializeRoutes(app *fiber.App, userMiddleware fiber.Handler) fiber.Router {
	users := app.Group("/user");
	usersById := users.Group("/:id")
	
	users.Use(userMiddleware)
	
	users.Get("/", c.GetAllUsers)

	usersById.Get("/", c.GetUser)
	usersById.Patch("/", c.UpdateUser)
	usersById.Delete("/", c.DeleteUser)

	return usersById
}