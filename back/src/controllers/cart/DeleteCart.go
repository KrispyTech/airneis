package cart

import (
	"github.com/gofiber/fiber/v2"
)

func DeleteProductFromCart(ctx *fiber.Ctx) error {
	_, errGetCart := GetCartOfLoggedUserIfExistsOrCreateOne(ctx)

	if errGetCart != nil {
		return errGetCart
	}

	// TODO: DELETE PRODUCT FROM CART AND RETURN CART
	return nil
}
