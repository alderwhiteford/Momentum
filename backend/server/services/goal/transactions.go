package goalService

import (
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/google/uuid"
)

func GetAllGoalsInDB(db *storage.PostgresDB) ([]Goal, error) {
	goals := []Goal{}
	if err := storage.InitializeSelectionOnAll(db, "goal", goals); err != nil {
		return nil, err
	}

	return goals, nil
}

func GetGoalByIDInDB(db *storage.PostgresDB, id uuid.UUID) (*Goal, error) {
	var goal Goal
	if err := storage.InitializeSingleSelectionOnEntity(db, "goal", goal, id); err != nil {
		return nil, err
	}

	return &goal, nil
}

func UpdateGoalInDB(db *storage.PostgresDB, goal UpdateGoal, id uuid.UUID) error {
	updateQuery, args, err := storage.BuildUpdateQuery("goal", id, goal)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	if err = storage.InitializeMutation(db, *updateQuery, args); err != nil {
		return err
	}

	return nil
}

func DeleteGoalInDB(db *storage.PostgresDB, id uuid.UUID) error {	
	if err := storage.InitializeDeletion(db, "goal", id); err != nil {
		return utilities.BadRequest(err.Error())
	}

	return nil
}


