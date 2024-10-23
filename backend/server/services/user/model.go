package userService

import "momentum/utilities"

type Provider string

const (
	Google Provider = "google"
)

type UserBaseModel struct {
	Provider Provider `json:"provider" db:"provider"`
	Email    string    `json:"email" db:"email" validate:"required"`
	Name     string    `json:"name" db:"name" validate:"required"`
	TimeZone string    `json:"timezone" validate:"timezone"`
}

type User struct {
	utilities.BaseModel
	UserBaseModel
}

type UpdateUser struct {
	Name     string `json:"name,omitempty" db:"name"`
	TimeZone string `json:"timezone,omitempty" db:"timezone" validate:"timezone"`
}
