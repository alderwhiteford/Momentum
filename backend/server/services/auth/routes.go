package authService

import "github.com/gofiber/fiber/v2"

func (c *AuthServiceImpl) InitializeRoutes(app *fiber.App) {
	auth := app.Group("/auth");
	
	auth.Post("/signin/google", c.SignIn)
}