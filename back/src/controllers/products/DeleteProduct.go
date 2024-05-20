package products

import (
	"gorm.io/gorm/clause"

	"github.com/KrispyTech/airneis/lib/shared/helpers"

	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteProduct(ctx *fiber.Ctx) error {
	var product model.Product

	if err := SelectProductByID(&product, ctx); err != nil {
		return helpers.SendNotFound(ctx, err)
	}

	if err := config.Database.Save(&product).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	productID := ctx.Params("productID")

	if err := config.Database.Clauses(clause.Returning{}).Delete(&product, productID).Error; err != nil {
		return helpers.SendBadRequest(ctx, err)
	}

	return sendProduct(ctx, product)
}
