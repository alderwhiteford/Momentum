package user

import "github.com/gofiber/fiber/v2"

func (c *UserServiceImpl) InitializeRoutes(app *fiber.App) {
	users := app.Group("/users");
	
	users.Get("/", c.GetAllUsers)
}