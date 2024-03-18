package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetCategoryByID(ctx *fiber.Ctx) error {
	var category model.Category

	if err := SelectCategoryByID(&category, ctx); err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return sendCategory(ctx, category)
}
