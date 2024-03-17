package controllers

import (
	"encoding/json"

	"strconv"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(ctx *fiber.Ctx) error {
	page, errParsing := strconv.Atoi(ctx.Query(constants.Page, constants.DefaultPageValue))
	if errParsing != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	var categories []model.Category
	errQuery := config.Database.Limit(constants.PaginationLimit).Offset((page - 1) * constants.PaginationLimit).Order("order_of_display asc").Find(&categories).Error
	if errQuery != nil {
		ctx.Status(constants.BadRequestStatus)
		return ctx.SendString(constants.BadRequestMessage)
	}

	jsonCategories, errMarshal := json.Marshal(categories)
	if errMarshal != nil {
		ctx.Status(constants.InternalServerErrorStatus)
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return ctx.Send(jsonCategories)
}
