package utilities

import "github.com/google/uuid"

type BaseModel struct {
	Id        uuid.UUID `json:"id" db:"id"`
	CreatedAt string    `json:"created_at" db:"created_at"`
}
