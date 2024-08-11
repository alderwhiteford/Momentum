package utilities

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type APIError struct {
	StatusCode int   	`json:"status_code"`
	Message    string 	`json:"message"` 
}

func NewAPIError(statusCode int, message string) APIError {
	return APIError{
		StatusCode: statusCode,
		Message: message,
	}
}

/** Extend the error API */
func (e APIError) Error() string {
	return fmt.Sprintf("api error: %d %v", e.StatusCode, e.Message)
}

func InternalServerError() APIError {
	return NewAPIError(500, "Internal Server Error")
}

func ErrorHandler(ctx *fiber.Ctx, err error) error {
	var apiError APIError

	if castedErr, ok := err.(APIError); ok {
		apiError = castedErr
	} else {
		apiError = InternalServerError()
	}

	return ctx.Status(apiError.StatusCode).JSON(apiError)
}