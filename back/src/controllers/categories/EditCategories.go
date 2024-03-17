package controllers

import (
	"encoding/json"

	"strconv"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditCategory(ctx *fiber.Ctx) error {
	categoryID, errConvertToString := strconv.Atoi(ctx.Params("categoryID"))
	if errConvertToString != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	var category model.Category
	if errGet := config.Database.First(&category, categoryID).Error; errGet != nil {
		ctx.Status(constants.NotFoundStatus)
		return ctx.SendString(constants.NotFoundMessage)
	}

	if errParsing := ctx.BodyParser(&category); errParsing != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	if err := config.Database.Save(&category).Error; err != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	jsonCategory, errMarshal := json.Marshal(category)
	if errMarshal != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategory)
}
