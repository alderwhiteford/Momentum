package milestoneService

import (
	"momentum/server/storage"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type MilestoneService interface { 
	InitializeRoutes(app fiber.Router, adminMiddleware fiber.Handler)
	GetMilestones(ctx *fiber.Ctx) error
	GetMilestone(ctx *fiber.Ctx) error
	UpdateMilestone(ctx *fiber.Ctx) error
	DeleteMilestone(ctx *fiber.Ctx) error
}

type MilestoneServiceImpl struct {
	db *storage.PostgresDB
	validate *validator.Validate
}

func NewMilestoneService(db *storage.PostgresDB, validate *validator.Validate) MilestoneService {
	return &MilestoneServiceImpl{db, validate}
}

func (c *MilestoneServiceImpl) GetMilestones(ctx *fiber.Ctx) error {
	return nil
}

func (c *MilestoneServiceImpl) GetMilestone(ctx *fiber.Ctx) error {
	return nil
}

func (c *MilestoneServiceImpl) UpdateMilestone(ctx *fiber.Ctx) error {
	return nil
}

func (c *MilestoneServiceImpl) DeleteMilestone(ctx *fiber.Ctx) error {
	return nil
}