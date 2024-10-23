package goalService

import (
	"momentum/utilities"

	"github.com/google/uuid"
)

type SuccessIndicator string
const (
	Completion SuccessIndicator = "COMPLETION"
	Target     SuccessIndicator = "TARGET"
	Satisfied  SuccessIndicator = "SATISFIED"
)

type GoalBaseModel struct {
	UserId                uuid.UUID        `json:"user_id" db:"user_id"`
	CompletedAt           *string          `json:"completed_at" db:"completed_at"`
	Description           string           `json:"description" db:"description"`
	EstimatedCompletionAt string           `json:"estimated_completion_at" db:"estimated_completion_at"`
	TheWhy                string           `json:"the_why" db:"the_why"`
	WhenSuccess           SuccessIndicator `json:"when_success" db:"when_success"`
}

type Goal struct {
	utilities.BaseModel
	GoalBaseModel
}
