package userService

import (
	goalService "momentum/server/services/goal"
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/google/uuid"
)

func GetAllUsersFromDB(db *storage.PostgresDB) ([]User, error) {
	users := []User{}
	if err := storage.InitializeSelectionOnAll(db, "users", users); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserByIDFromDB(db *storage.PostgresDB, id uuid.UUID) (*User, error) {
	var user User
	if err := storage.InitializeSingleSelectionOnEntity(db, "users", user, id); err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUserInDB(db *storage.PostgresDB, user User) error {
	query, _ := storage.BuildCreateQuery("users", user)

	if err := storage.InitializeCreation(db, *query, user); err != nil {
		return err
	}

	return nil
}

func UpdateUserInDB(db *storage.PostgresDB, id uuid.UUID, updateUser UpdateUser) error {
	updateQuery, args, err := storage.BuildUpdateQuery("users", id, updateUser)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	if err := storage.InitializeMutation(db, *updateQuery, args); err != nil {
		return err
	}

	return nil
}

func DeleteUserInDB(db *storage.PostgresDB, id uuid.UUID) error {
	if err := storage.InitializeDeletion(db, "users", id); err != nil {
		return err
	}
	
	return nil
}

func GetUserGoalsInDB(db *storage.PostgresDB, id uuid.UUID) ([]goalService.Goal, error) {
	goals := []goalService.Goal{}
	if err := storage.InitializeMultiSelectionOnEntity(db, "goal", goals, id, "user_id"); err != nil {
		return nil, err
	}

	return goals, nil
}

func CreateUserGoalInDB(db *storage.PostgresDB, goal goalService.Goal) error {
	query, _ := storage.BuildCreateQuery("goal", goal)

	if err := storage.InitializeCreation(db, *query, goal); err != nil {
		return err
	}

	return nil
}
