package cart

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/gofiber/fiber/v2"
)

func GetCart(ctx *fiber.Ctx) error {
	_, err := GetCartOfLoggedUser(ctx)
	if err != nil {
		return helpers.SendNotFound(ctx, err)
	}

	//TODO: SEND CART AFTER PR MERGED | rename _ to cart
	return nil
}
