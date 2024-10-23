package userService

import (
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService interface { 
	InitializeRoutes(app *fiber.App, middleware fiber.Handler) fiber.Router

	GetAllUsers(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type UserServiceImpl struct {
	db *storage.PostgresDB
	validator *validator.Validate
}

func NewUserService(db *storage.PostgresDB, validator *validator.Validate) UserService {
	return &UserServiceImpl{db, validator}
}

// Fetch all users:
func (c *UserServiceImpl) GetAllUsers(ctx *fiber.Ctx) error {
	users, err := GetAllUsersFromDB(c.db);
	if err != nil {
		return utilities.InternalServerError(err.Error());
	}

	return ctx.Status(fiber.StatusOK).JSON(users);
}

// Retrieve a user by their id:
func (c *UserServiceImpl) GetUser(ctx *fiber.Ctx) error {
	// Extract the user_id from the path:
	pathUserID := ctx.Params("id")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}

	user, err := GetUserByIDFromDB(c.db, uuid)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(&user)
}

// Update a user:
func (c *UserServiceImpl) UpdateUser(ctx *fiber.Ctx) error {
	// Extract the user_id from the path:
	pathUserID := ctx.Params("id")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}
	
	var updateRequestBody UpdateUser
	if err := ctx.BodyParser(&updateRequestBody); err != nil {
		return utilities.BadRequest("failed to parse request body")
	}

	err = utilities.Validate(c.validator, updateRequestBody)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	if err = UpdateUserInDB(c.db, uuid, updateRequestBody); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}

// Delete a user:
func (c *UserServiceImpl) DeleteUser(ctx *fiber.Ctx) error {
	pathUserID := ctx.Params("id")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}

	if err = DeleteUserInDB(c.db, uuid); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusOK)
}
