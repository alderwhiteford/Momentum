package userService

import (
	"github.com/gofiber/fiber/v2"
)

func (c *UserServiceImpl) InitializeRoutes(app *fiber.App, middleware fiber.Handler) {
	users := app.Group("/users");
	users.Use(middleware)
	
	users.Get("/", c.GetAllUsers)
}