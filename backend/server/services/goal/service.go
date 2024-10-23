package goalService

import (
	"momentum/server/storage"

	"github.com/gofiber/fiber/v2"
)

type GoalService interface { 
	InitializeRoutes(router fiber.Router)
	CreateGoal(ctx *fiber.Ctx) error
}

type GoalServiceImpl struct {
	db *storage.PostgresDB
}

func NewGoalService(db *storage.PostgresDB) *GoalServiceImpl {
	return &GoalServiceImpl{db}
}

func (c *GoalServiceImpl) CreateGoal(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK);
}
