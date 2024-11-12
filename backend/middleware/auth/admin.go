package middleware

import (
	"github.com/gofiber/fiber/v2"
)

func NewAdmin() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Set the locals to check if user is an admin:
		isAdmin, ok := ctx.Locals("is_admin").(bool)
		if !ok || !isAdmin {
			return ctx.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}

		return ctx.Next()
	}
}
