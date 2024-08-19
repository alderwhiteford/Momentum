package userService

import (
	"momentum/server/storage"
)

func GetAllUsersFromDB(db *storage.PostgresDB) ([]User, error) {
	users := []User{}
	err := db.Select(&users, "SELECT * FROM users")

	if err != nil {
		return nil, err
	}

	return users, nil
}

func GetUserByIDFromDB(db *storage.PostgresDB, id string) (*User, error) {
	var user User
	err := db.Get(&user, "SELECT * FROM users WHERE user_id = $1", id)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func CreateUserInDB(db *storage.PostgresDB, user User) error {
	_, err := db.NamedExec("INSERT INTO users (user_id, provider, email, name) VALUES (:user_id, :provider, :email, :name)", user)
	if err != nil {
		return err
	}

	return nil
}
