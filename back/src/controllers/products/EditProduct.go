package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func EditProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if errGet := SelectProductByID(&product, ctx); errGet != nil {
		return helpers.SendNotFound(ctx)
	}

	if errParsing := ctx.BodyParser(&product); errParsing != nil {
		return helpers.SendBadRequest(ctx)
	}

	if err := config.Database.Save(&product).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return sendProduct(ctx, product)
}
