package auth

import (
	"fmt"
	userService "momentum/server/services/user"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func FromRequest(ctx *fiber.Ctx, settings utilities.AuthSettings) (*jwt.Token, error) {
	// Read JWT token from the headers:
	headers := ctx.GetReqHeaders()
	authorizationHeaders := headers["Authorization"]
	if len(authorizationHeaders) == 0 {
		return nil, fmt.Errorf("no token provided")
	}

	token := authorizationHeaders[0]
	if token == "" {
		return nil, fmt.Errorf("no token provided")
	}

	// Verify the token:
	jwtToken, err := verifyToken(token, settings.JwtSecret)
	if err != nil {
		return nil, fmt.Errorf("invalid token")
	}

	return jwtToken, nil
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

func GetUserIDFromClaims(claims jwt.Claims) (string, error) {
	claims, ok := claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("failed to parse claims")
	}

	userID, err := claims.GetSubject()
	if err != nil {
		return "", fmt.Errorf("failed to parse user ID")
	}

	return userID, nil
}

func IntoUser(token jwt.Token) (*userService.User, error) {
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

	return &userService.User{
		ID: userID,
		Provider: userService.Provider(provider),
		Email: email,
		Name: name,
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


