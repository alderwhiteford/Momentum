package authService

import "github.com/gofiber/fiber/v2"

func (c *AuthServiceImpl) InitializeRoutes(app *fiber.App) {
	users := app.Group("/auth");
	
	users.Post("/signin/google", c.SignIn)
}