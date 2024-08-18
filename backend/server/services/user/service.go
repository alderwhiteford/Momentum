package userService

import (
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/gofiber/fiber/v2"
)

type UserService interface { 
	InitializeRoutes(app *fiber.App, middleware fiber.Handler)
	GetAllUsers(ctx *fiber.Ctx) error
}

type UserServiceImpl struct {
	db *storage.PostgresDB
}

func NewUserService(db *storage.PostgresDB) UserService {
	return &UserServiceImpl{db}
}

func (c *UserServiceImpl) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := GetAllUsersFromDB(c.db);
	if err != nil {
		return utilities.InternalServerError(err.Error());
	}

	return ctx.Status(200).JSON(users);
}
