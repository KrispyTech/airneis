package controllers

import (
	"encoding/json"
	"github.com/KrispyTech/airneis/config"
	"github.com/KrispyTech/airneis/lib/shared/constants"
	model "github.com/KrispyTech/airneis/src/db/models"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

func EditCategory(ctx *fiber.Ctx) error {
	categoryID, err := strconv.Atoi(ctx.Params("categoryID"))
	if err != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	var category model.Category
	config.Database.First(&category, categoryID)

	if errParsing := ctx.BodyParser(&category); errParsing != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	config.Database.Save(&category)

	jsonCategory, err := json.Marshal(category)
	if err != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategory)
}
