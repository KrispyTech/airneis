package controllers

import (
	"encoding/json"
	"github.com/KrispyTech/airneis/config"
	"github.com/KrispyTech/airneis/lib/shared/constants"
	model "github.com/KrispyTech/airneis/src/db/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm/clause"
)

func DeleteCategory(ctx *fiber.Ctx) error {
	var category model.Category
	categoryId := ctx.Params("categoryID")

	// Set the order of display to nil, so it become available for another as there is a unique constraint
	if err := config.Database.First(&category, categoryId).Error; err != nil {
		ctx.Status(constants.NotFoundStatus)
		return ctx.SendString(constants.NotFoundMessage)
	}

	category.OrderOfDisplay = nil
	if err := config.Database.Save(&category).Error; err != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	if err := config.Database.Clauses(clause.Returning{}).Delete(&category, categoryId).Error; err != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	jsonCategory, err := json.Marshal(category)
	if err != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategory)
}
