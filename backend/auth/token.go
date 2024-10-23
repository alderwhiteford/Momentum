package auth

import (
	"fmt"
	userService "momentum/server/services/user"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func FromRequest(ctx *fiber.Ctx, settings utilities.AuthSettings) (*jwt.Token, bool, error) {
	// Read JWT token from the headers:
	headers := ctx.GetReqHeaders()
	authorizationHeaders := headers["Authorization"]
	if len(authorizationHeaders) == 0 {
		return nil, false, fmt.Errorf("no token provided")
	}

	token := authorizationHeaders[0]
	if token == "" {
		return nil, false, fmt.Errorf("no token provided")
	}

	// Check to see if the token is the admin:
	if token == settings.AdminToken {
		return nil, true, nil
	}

	// Verify the token:
	jwtToken, err := verifyToken(token, settings.JwtSecret)
	if err != nil {
		return nil, false, fmt.Errorf("invalid token")
	}

	return jwtToken, false, nil
}

func GetRoleFromToken(token jwt.Token) (string, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse claims")
	}

	role, ok := claims["role"].(string)
	if !ok {
		return "", fmt.Errorf("failed to parse role")
	}

	return role, nil
}

func GetUserIDFromClaims(claims jwt.Claims) (*uuid.UUID, error) {
	claims, ok := claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID")
	}

	userIDAsUUID, err := uuid.Parse(userID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID")
	}

	return &userIDAsUUID, nil
}

func IntoProviderUser(token jwt.Token) (*userService.User, error) {
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	// Retrieve the user ID:
	userID, err := GetUserIDFromClaims(claims)
	if err != nil {
		return nil, err
	}

	userMetadata, ok := claims["user_metadata"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to parse user metadata")
	}

	// Retrieve user information:
	email, ok := userMetadata["email"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to parse email")
	}

	name, ok := userMetadata["full_name"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to parse email")
	}

	// Retrieve provider information:
	appMetadata, ok := claims["app_metadata"].(map[string]interface{})
	if !ok {
		return nil, fmt.Errorf("failed to parse provider")
	}

	provider, ok := appMetadata["provider"].(string)
	if !ok {
		return nil, fmt.Errorf("failed to parse provider")
	}

	parsedProvider := userService.Provider(provider)

	return &userService.User{
		BaseModel: utilities.BaseModel{
			Id: *userID,
		},
		UserBaseModel: userService.UserBaseModel{
			Provider: parsedProvider,
			Email: email,
			Name: name,
		},
	}, nil
}

func verifyToken(token string, secret string) (*jwt.Token, error) {
	// Parse the token:
	jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		// Confirm the signing method:
		if token.Method.Alg() != "HS256" {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Method.Alg())
		}

		return []byte(secret), nil
	})

	// Check for any errors:
	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %v", err)
	}

	// Check if the token is valid / not expired:
	if !jwtToken.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return jwtToken, nil
}
