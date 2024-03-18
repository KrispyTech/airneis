package controllers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func sendCategory(ctx *fiber.Ctx, category model.Category) error {
	categoryJSON, err := json.Marshal(category)
	if err != nil {
		ctx.Status(constants.InternalServerErrorStatus)

		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(categoryJSON)
}

func sendCategories(ctx *fiber.Ctx, categories []model.Category) error {
	jsonCategories, errMarshal := json.Marshal(categories)
	if errMarshal != nil {
		ctx.Status(constants.InternalServerErrorStatus)

		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategories)
}
