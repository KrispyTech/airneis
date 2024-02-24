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
	config.Database.Clauses(clause.Returning{}).Delete(&category, categoryId)

	jsonCategory, err := json.Marshal(category)
	if err != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategory)
}
