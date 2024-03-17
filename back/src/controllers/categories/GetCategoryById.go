package controllers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetCategoryByID(ctx *fiber.Ctx) error {
	var category model.Category
	categoryID := ctx.Params("categoryID")

	if err := config.Database.First(&category, categoryID).Error; err != nil {
		ctx.Status(constants.NotFoundStatus)
		return ctx.SendString(constants.NotFoundMessage)
	}

	jsoncCategory, err := json.Marshal(category)
	if err != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsoncCategory)
}
