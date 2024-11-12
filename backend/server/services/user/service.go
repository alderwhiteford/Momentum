package userService

import (
	goalService "momentum/server/services/goal"
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type UserService interface { 
	InitializeRoutes(app *fiber.App, userMiddleware fiber.Handler, adminMiddleware fiber.Handler)

	// ENTITY SPECIFIC
	GetAllUsers(ctx *fiber.Ctx) error
	GetUser(ctx *fiber.Ctx) error
	UpdateUser(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error

	// RELATIONSHIPS
	GetUserGoals(ctx *fiber.Ctx) error
	CreateUserGoal(ctx *fiber.Ctx) error
}

type UserServiceImpl struct {
	db *storage.PostgresDB
	validate *validator.Validate
}

func NewUserService(db *storage.PostgresDB, validate *validator.Validate) UserService {
	return &UserServiceImpl{db, validate}
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
	pathUserID := ctx.Params("userId")

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
	pathUserID := ctx.Params("userId")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}
	
	var updateRequestBody UpdateUser
	if err := ctx.BodyParser(&updateRequestBody); err != nil {
		return utilities.BadRequest("failed to parse request body")
	}

	err = utilities.Validate(c.validate, updateRequestBody)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	if err = UpdateUserInDB(c.db, uuid, updateRequestBody); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// Delete a user:
func (c *UserServiceImpl) DeleteUser(ctx *fiber.Ctx) error {
	pathUserID := ctx.Params("userId")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}

	if err = DeleteUserInDB(c.db, uuid); err != nil {
		return err
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

// Get user goals:
func (c *UserServiceImpl) GetUserGoals(ctx *fiber.Ctx) error {
	pathUserID := ctx.Params("userId")

	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}

	goals, err := GetUserGoalsInDB(c.db, uuid)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(goals)
}

// Create user goal:
func (c *UserServiceImpl) CreateUserGoal(ctx *fiber.Ctx) error {
	pathUserID := ctx.Params("userId")
	
	// Parse the uuid:
	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest("invalid id")
	}
	
	// Extract the goal from the body
	var goal goalService.Goal
	if err := ctx.BodyParser(&goal); err != nil {
		return utilities.BadRequest("failed to parse request body")
	}

	// Validate the request body:
	err = utilities.Validate(c.validate, goal)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	// Add the user id to the goal struct
	goal.UserId = uuid

	if err = CreateUserGoalInDB(c.db, goal); err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated);
}
