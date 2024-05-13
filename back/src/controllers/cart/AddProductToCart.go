package cart

import (
	"github.com/gofiber/fiber/v2"
)

func AddProductToCart(ctx *fiber.Ctx) error {
	_, errGetCart := GetCartOfLoggedUserIfExistsOrCreateOne(ctx)

	if errGetCart != nil {
		return errGetCart
	}

	// TODO: ADD PRODUCT TO CART AND RETURN CART

	return nil
}
