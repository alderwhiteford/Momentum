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
	InitializeRoutes(router fiber.Router, adminMiddleware fiber.Handler)
	GetAllGoals(ctx *fiber.Ctx) error
	GetGoal(ctx *fiber.Ctx) error
	UpdateGoal(ctx *fiber.Ctx) error
	DeleteGoal(ctx *fiber.Ctx) error
}

type GoalServiceImpl struct {
	db *storage.PostgresDB
	validate *validator.Validate
}

func NewGoalService(db *storage.PostgresDB, validate *validator.Validate) GoalService {
	return &GoalServiceImpl{db, validate}
}

func (c *GoalServiceImpl) GetAllGoals(ctx *fiber.Ctx) error {
	goals, err := GetAllGoalsInDB(c.db)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(goals)
}

func (c *GoalServiceImpl) GetGoal(ctx *fiber.Ctx) error {
	pathGoalId := ctx.Params("goalId")

	// Parse the uuid:
	uuid, err := uuid.Parse(pathGoalId)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed to parse id: %s", pathGoalId))
	}

	goal, err := GetGoalByIDInDB(c.db, uuid)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.Status(fiber.StatusOK).JSON(goal)
}

func (c *GoalServiceImpl) UpdateGoal(ctx *fiber.Ctx) error {
	// Extract the goal from the body
	var goal UpdateGoal
	if err := ctx.BodyParser(&goal); err != nil {
		return utilities.BadRequest("failed to parse request body")
	}

	// Validate the request body:
	err := utilities.Validate(c.validate, goal)
	if err != nil {
		return utilities.BadRequest(err.Error())
	}

	pathGoalId := ctx.Params("goalId")
	
	// Parse the uuid:
	uuid, err := uuid.Parse(pathGoalId)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed to parse id: %s", pathGoalId))
	}

	if err := UpdateGoalInDB(c.db, goal, uuid); err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}

func (c *GoalServiceImpl) DeleteGoal(ctx *fiber.Ctx) error {
	pathGoalId := ctx.Params("goalId")

	// Parse the uuid:
	uuid, err := uuid.Parse(pathGoalId)
	if err != nil {
		return utilities.BadRequest(fmt.Sprintf("failed to parse id: %s", pathGoalId))
	}

	if err := DeleteGoalInDB(c.db, uuid); err != nil {
		return utilities.BadRequest(err.Error())
	}

	return ctx.SendStatus(fiber.StatusNoContent)
}
