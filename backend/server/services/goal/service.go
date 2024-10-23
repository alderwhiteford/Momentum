package goalService

import (
	"fmt"
	"momentum/server/storage"
	"momentum/utilities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type GoalService interface { 
	InitializeRoutes(router fiber.Router)
	CreateGoal(ctx *fiber.Ctx) error
}

type GoalServiceImpl struct {
	db *storage.PostgresDB
	validate *validator.Validate
}

func NewGoalService(db *storage.PostgresDB, validate *validator.Validate) *GoalServiceImpl {
	return &GoalServiceImpl{db, validate}
}

func (c *GoalServiceImpl) CreateGoal(ctx *fiber.Ctx) error {
	// Extract the goal from the body
	var goal GoalBaseModel
	if err := ctx.BodyParser(&goal); err != nil {
		return utilities.BadRequest("failed to parse request body")
	}

	// Validate the request body:
	err := utilities.Validate(c.validate, goal)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	pathUserID := ctx.Params("id")
	
	// Parse the uuid:
	uuid, err := uuid.Parse(pathUserID)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed to parse id: %s", pathUserID))
	}

	// Add the user id to the goal struct
	goal.UserId = uuid

	if err = CreateGoalInDB(c.db, goal); err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.SendStatus(fiber.StatusCreated);
}
