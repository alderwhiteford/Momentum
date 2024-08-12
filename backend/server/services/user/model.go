package user

import (
	"github.com/google/uuid"
)

type User struct {
	ID 			uuid.UUID 	`json:"user_id" db:"user_id"`
	CreatedAt 	string 		`json:"created_at" db:"created_at"`
	FirstName 	string		`json:"first_name" db:"first_name"`
	LastName 	string 		`json:"last_name" db:"last_name"`
}