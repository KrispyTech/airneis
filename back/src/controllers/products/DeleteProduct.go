package controllers

import (
	"gorm.io/gorm/clause"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if err := SelectProductByID(&product, ctx); err != nil {
		return helpers.SetStatusAndMessages(
			constants.NotFoundStatus,
			constants.NotFoundMessage,
			ctx)
	}

	if err := config.Database.Save(&product).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	productID := ctx.Params("productID")

	if err := config.Database.Clauses(clause.Returning{}).Delete(&product, productID).Error; err != nil {
		return helpers.SetStatusAndMessages(
			constants.BadRequestStatus,
			constants.BadRequestMessage,
			ctx)
	}

	return sendProduct(ctx, product)
}
