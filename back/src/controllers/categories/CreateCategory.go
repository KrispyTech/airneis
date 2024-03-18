package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	"github.com/KrispyTech/airneis/pkg/validator"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateCategory(ctx *fiber.Ctx) error {
	var category model.Category

	if err := ctx.BodyParser(&category); err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	if err := validator.ValidateInput(category); err != nil {
		return helpers.SetStatusAndMessages(constants.BadRequestStatus, err.Error(), ctx)
	}

	if err := config.Database.Create(&category).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendCategory(ctx, category)
}
