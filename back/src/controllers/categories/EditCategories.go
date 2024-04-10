package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditCategory(ctx *fiber.Ctx) error {
	var category model.Category

	if errGet := SelectCategoryByID(&category, ctx); errGet != nil {
		return helpers.SendNotFound(ctx)
	}

	if errParsing := ctx.BodyParser(&category); errParsing != nil {
		return helpers.SendBadRequest(ctx)
	}

	if err := config.Database.Save(&category).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return sendCategory(ctx, category)
}
