package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditCategory(ctx *fiber.Ctx) error {
	var category model.Category

	if errGet := SelectCategoryByID(&category, ctx); errGet != nil {
		return helpers.SetStatusAndMessages(
			constants.NotFoundStatus,
			constants.NotFoundMessage,
			ctx)
	}

	if errParsing := ctx.BodyParser(&category); errParsing != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	if err := config.Database.Save(&category).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendCategory(ctx, category)
}
