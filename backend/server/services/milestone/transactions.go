package milestoneService

import (
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/google/uuid"
)

func GetAllMilestonesInDB(db *storage.PostgresDB) ([]Milestone, error) {
	var milestones []Milestone
	if err := storage.InitializeSelectionOnAll(db, "milestone", milestones); err != nil {
		return nil, err
	}

	return milestones, nil
}

func GetMilestoneByIDInDB(db *storage.PostgresDB, id uuid.UUID) (*Milestone, error) {
	var milestone Milestone
	if err := storage.InitializeSingleSelectionOnEntity(db, "milestone", milestone, id); err != nil {
		return nil, err
	}

	return &milestone, nil
}

func UpdateMilestoneInDB(db *storage.PostgresDB, milestone UpdateMilestone, id uuid.UUID) error {
	updateQuery, args, err := storage.BuildUpdateQuery("milestone", id, milestone)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	if err = storage.InitializeMutation(db, *updateQuery, args); err != nil {
		return err
	}

	return nil
}