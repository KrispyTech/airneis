package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func CreateProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if err := ctx.BodyParser(&product); err != nil {
		return helpers.SendBadRequest(ctx)
	}

	if err := config.Database.Create(&product).Error; err != nil {
		return helpers.SendInternalError(ctx)
	}

	return sendProduct(ctx, product)
}
