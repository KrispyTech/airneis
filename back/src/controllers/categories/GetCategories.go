package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func GetCategories(ctx *fiber.Ctx) error {
	// We need to convert the page because query category must be an int
	page, errParsing := helpers.ConvertStringToInt(ctx.Query(constants.Page, constants.DefaultPageValue))
	if errParsing != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	categories, err := queryCategories(page, ctx)
	if err != nil {
		return ctx.SendString(constants.InternalServerErrorMessage)
	}

	return sendCategories(ctx, categories)
}

func queryCategories(page int, ctx *fiber.Ctx) ([]model.Category, error) {
	pagLimit := constants.PaginationLimit
	offset := (page - 1) * pagLimit
	db := config.Database

	cats := make([]model.Category, pagLimit)
	err := db.Limit(pagLimit).Preload("Products").Offset(offset).Order("order_of_display asc").Find(&cats)
	if err.Error != nil {
		ctx.Status(constants.BadRequestStatus)

		return cats, ctx.SendString(constants.BadRequestMessage)
	}

	return cats, nil
}
