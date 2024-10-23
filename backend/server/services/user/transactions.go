package userService

import (
	"fmt"
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/google/uuid"
)

func GetAllUsersFromDB(db *storage.PostgresDB) ([]User, error) {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByIDFromDB(db *storage.PostgresDB, id uuid.UUID) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserInDB(db *storage.PostgresDB, user User) error {
	_, err := db.NamedExec("INSERT INTO users (id, provider, email, name) VALUES (:id, :provider, :email, :name)", user)
	if err != nil {
		return err
	}

	return nil
}

func UpdateUserInDB(db *storage.PostgresDB, id uuid.UUID, updateUser UpdateUser) error {
	updateQuery, args, err := utilities.BuildUpdateQuery("users", id, updateUser)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	result, err := db.Exec(*updateQuery, args...)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("error executing update: %s", err.Error()))
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return utilities.InternalServerError(fmt.Sprintf("something went wrong: %s", err.Error()))
	}
	if rows == 0 {
		return utilities.BadRequest(fmt.Sprintf("failed to find user in db with id: %s", id))
	}

	return nil
}

func DeleteUserInDB(db *storage.PostgresDB, id uuid.UUID) error {
	result, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("error executing delete: %s", err.Error()))
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return utilities.InternalServerError(fmt.Sprintf("something went wrong: %s", err.Error()))
	}
	if rows == 0 {
		return utilities.BadRequest(fmt.Sprintf("user does not exist: %s", id))
	}

	return nil
}
