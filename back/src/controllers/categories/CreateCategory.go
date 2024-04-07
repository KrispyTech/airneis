package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(ctx *fiber.Ctx) error {
	var category model.Category

	if err := ctx.BodyParser(&category); err != nil {
		return helpers.SendBadRequest(ctx)
	}

	if err := config.Database.Create(&category).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return sendCategory(ctx, category)
}
