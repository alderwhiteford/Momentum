package user

import (
	"momentum/server/storage"
)

func getAllUsersFromDB(db *storage.PostgresDB) ([]User, error) {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}
