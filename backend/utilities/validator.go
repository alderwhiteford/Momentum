package utilities

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

func NewValidator() *validator.Validate {
	validator := validator.New(validator.WithRequiredStructEnabled())

	validator.RegisterValidation("timezone", validateTimeZone)

	return validator
}

func Validate(validate *validator.Validate, toValidate interface{}) error {
	err := validate.Struct(toValidate)
	if err != nil {
		return fmt.Errorf("error: %s", err.Error())
	}

	return nil
}

func validateTimeZone(fl validator.FieldLevel) bool {
	_, err := time.LoadLocation(fl.Field().String())
	return err == nil
}

