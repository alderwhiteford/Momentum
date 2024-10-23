package goalService

import (
	"momentum/server/storage"
	"momentum/utilities"
)

func CreateGoalInDB(db *storage.PostgresDB, goal GoalBaseModel) error {
	query, _ := utilities.BuildCreateQuery("goal", goal)

	_, err := db.NamedExec(*query, goal)
	if err != nil {
		return err
	}

	return nil
}
