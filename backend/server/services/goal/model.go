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
	UserId                uuid.UUID        `json:"user_id,omitempty" db:"user_id"`
	CompletedAt           *string          `json:"completed_at,omitempty" db:"completed_at"`
	Description           string           `json:"description" db:"description" validate:"required"`
	EstimatedCompletionAt string           `json:"estimated_completion_at" db:"estimated_completion_at" validate:"required"`
	TheWhy                string           `json:"the_why" db:"the_why" validate:"required"`
	WhenSuccess           SuccessIndicator `json:"when_success" db:"when_success" validate:"required,oneof=COMPLETION TARGET SATISFIED"`
}

type UpdateGoal struct {
	CompletedAt           *string          `json:"completed_at,omitempty" db:"completed_at"`
	Description           string           `json:"description" db:"description" validate:"required"`
	EstimatedCompletionAt string           `json:"estimated_completion_at" db:"estimated_completion_at" validate:"required"`
	TheWhy                string           `json:"the_why" db:"the_why" validate:"required"`
	WhenSuccess           SuccessIndicator `json:"when_success" db:"when_success" validate:"required,oneof=COMPLETION TARGET SATISFIED"`
}

type Goal struct {
	utilities.BaseModel
	GoalBaseModel
}
