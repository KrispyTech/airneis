package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

type InputData struct {
	CategoryID uint `json:"categoryID"`
}

func EditProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if errGet := SelectProductByID(&product, ctx); errGet != nil {
		return helpers.SendNotFound(ctx)
	}

	if errParsing := ctx.BodyParser(&product); errParsing != nil {
		return helpers.SendBadRequest(ctx)
	}

	if err := setCategoryIfModifyied(ctx, &product); err != nil {
		return helpers.SendInternalError(ctx)
	}

	if err := config.Database.Save(&product).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return sendProduct(ctx, product)
}

func setCategoryIfModifyied(ctx *fiber.Ctx, product *model.Product) error {
	var inputData InputData

	if errParsing := ctx.BodyParser(&inputData); errParsing != nil {
		return errParsing
	}

	if inputData.CategoryID == 0 {
		return nil
	}

	var category model.Category

	if err := config.Database.Preload("Products").First(&category, inputData.CategoryID).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	product.Category = category

	return nil
}
