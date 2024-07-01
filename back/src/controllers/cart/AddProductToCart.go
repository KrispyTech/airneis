package cart

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/src/controllers/products"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/gofiber/fiber/v2"
)

func AddProductToCart(ctx *fiber.Ctx) error {
	cart, errGetCart := GetCartOfLoggedUserIfExistsOrCreateOne(ctx)

	if errGetCart != nil {
		return errGetCart
	}

	var product model.Product

	if err := products.SelectProductByID(&product, ctx); err != nil {
		return helpers.SendNotFound(ctx, err)
	}

	cart.Products = append(cart.Products, model.OrderProduct{Product: product, Quantity: 1})

	// TODO: RETURN CART WHEN AUTH PR IS MERGED
	return nil
}
