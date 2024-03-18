package controllers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/pkg/config"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func sendCategory(ctx *fiber.Ctx, category model.Category) error {
	categoryJSON, err := json.Marshal(category)
	if err != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(categoryJSON)
}

func sendCategories(ctx *fiber.Ctx, categories []model.Category) error {
	jsonCategories, errMarshal := json.Marshal(categories)
	if errMarshal != nil {
		return helpers.SetStatusAndMessages(
			constants.InternalServerErrorStatus,
			constants.InternalServerErrorMessage,
			ctx)
	}

	return ctx.Send(jsonCategories)
}

func SelectCategoryByID(category *model.Category, ctx *fiber.Ctx) error {
	categoryID := ctx.Params("categoryID")

	return config.Database.First(&category, categoryID).Error
}
