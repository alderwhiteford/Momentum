package authService

import (
	"momentum/auth"
	userService "momentum/server/services/user"
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
)

type AuthService interface { 
	InitializeRoutes(app *fiber.App)
	SignIn(ctx *fiber.Ctx) error
}

type AuthServiceImpl struct {
	db *storage.PostgresDB
	authSettings utilities.AuthSettings
}

func NewAuthService(db *storage.PostgresDB, authSettings utilities.AuthSettings) AuthService {	
	return &AuthServiceImpl{
		db: db,
		authSettings: authSettings,
	}
}

func (c *AuthServiceImpl) SignIn(ctx *fiber.Ctx) error {
	// Read + parse JWT token from the headers:
	jwtToken, _, err := auth.FromRequest(ctx, c.authSettings)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	// Parse the token into a user:
	userFromToken, err := auth.IntoProviderUser(*jwtToken)
	if err != nil {
		return utilities.InternalServerError(err.Error())
	}

	// Check to see if the user exists:
	user, err := userService.GetUserByIDFromDB(c.db, userFromToken.Id)
	if err != nil {
		// Create a user if they don't exist:
		err = userService.CreateUserInDB(c.db, *userFromToken)
		if err != nil {
			return utilities.InternalServerError(err.Error())
		}
		user = userFromToken
	}

	return ctx.Status(200).JSON(user);
}
