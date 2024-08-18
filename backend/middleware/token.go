package middleware

import (
	"momentum/auth"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
)

func NewToken(settings utilities.AuthSettings) fiber.Handler {
	publicEndpoints := []string{
		"/auth/signin",
	}

	return func(ctx *fiber.Ctx) error {
		// Check if the current path is a public endpoint:
		for _, endpoint := range publicEndpoints {
			if ctx.Path() == endpoint {
				return ctx.Next()
			}
		}
	
		jwtToken, err := auth.FromRequest(ctx, settings)
		if err != nil {
			return utilities.Unauthorized(err.Error())
		}

		// Parse the role from the token:
		role, err := auth.GetRoleFromToken(*jwtToken)
		if err != nil {
			return utilities.InternalServerError(err.Error())
		}

		// Find the userID from context:
		userID := role
		if role == "authenticated" {
			userID, err = auth.GetUserIDFromClaims(jwtToken.Claims)
			if err != nil {
				return utilities.InternalServerError(err.Error())
			}
		}

		// Instantiate the userID in the context:
		ctx.Locals("user_id", userID)

		// Move to the next request:
		return ctx.Next()
	}
}