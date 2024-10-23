package middleware

import (
	"momentum/auth"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
)

func NewToken(settings utilities.AuthSettings) fiber.Handler {
	publicEndpoints := []string{
		"/auth/signin/email",
		"/auth/signup/email",
	}

	return func(ctx *fiber.Ctx) error {
		// Check if the current path is a public endpoint:
		for _, endpoint := range publicEndpoints {
			if ctx.Path() == endpoint {
				return ctx.Next()
			}
		}
	
		// Parse the token from the request
		jwtToken, isAdmin, err := auth.FromRequest(ctx, settings)
		if err != nil {
			return utilities.Unauthorized(err.Error())
		}

		// Set the locals to check if user is an admin:
		ctx.Locals("is_admin", isAdmin)
		if isAdmin {
			return ctx.Next()
		}

		// Parse the role from the token:
		role, err := auth.GetRoleFromToken(*jwtToken)
		if err != nil {
			return utilities.InternalServerError(err.Error())
		}

		// Check to see if the user is authenticated
		if role != "authenticated" {
			return utilities.Unauthorized("Unauthorized")
		}

		userID, err := auth.GetUserIDFromClaims(jwtToken.Claims)
		if err != nil {
			return utilities.InternalServerError(err.Error())
		}

		// Instantiate the userID in the context:
		ctx.Locals("user_id", userID)

		// Move to the next request:
		return ctx.Next()
	}
}