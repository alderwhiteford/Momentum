package milestoneService

import (
	"momentum/utilities"

	"github.com/google/uuid"
)

type MilestoneBaseModel struct {
	GoalId      uuid.UUID `json:"goal_id,omitempty" db:"goal_id"`
	CompletedAt string    `json:"completed_at,omitempty" db:"completed_at"`
	Title       string    `json:"title" db:"title" validate:"required"`
	Notes       string    `json:"notes" db:"notes" validate:"required"`
	Deadline    string    `json:"deadline" db:"deadline" validate:"required"`
}

type UpdateMilestone struct {
	GoalId      uuid.UUID `json:"goal_id,omitempty" db:"goal_id"`
	CompletedAt string    `json:"completed_at,omitempty" db:"completed_at"`
	Title       string    `json:"title,omitempty" db:"title"`
	Notes       string    `json:"notes,omitempty" db:"notes"`
	Deadline    string    `json:"deadline,omitempty" db:"deadline"`
}

type Milestone struct {
	utilities.BaseModel
	MilestoneBaseModel
}
