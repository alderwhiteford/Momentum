package middleware

import (
	"fmt"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
)

func NewUser(settings utilities.AuthSettings) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		// Check to see if the current user is an admin:
		isAdmin := ctx.Locals("is_admin")
		if isAdmin == true {
			return ctx.Next()
		}
		
		// Get the userID from the context:
		userID := ctx.Locals("user_id")
		if userID == nil {
			return utilities.Unauthorized("User not authenticated")
		}

		// Extract the user_id from the path:
		pathUserID := ctx.Params("id")

		// Check if the user is authorized to access the resource:
		if pathUserID != userID && userID != "service_role" {
			return utilities.Forbidden(fmt.Sprintf("User %s is not authorized to access this resource", userID))
		}

		return ctx.Next()
	}
}