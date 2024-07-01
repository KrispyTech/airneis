package cart

import "github.com/gofiber/fiber/v2"

func UpdateCartQuantity(ctx *fiber.Ctx) error {
	_, errGetCart := GetCartOfLoggedUserIfExistsOrCreateOne(ctx)

	if errGetCart != nil {
		return errGetCart
	}

	// TODO: UPDATE CART QUANTITY AND RETURN CART
	return nil
}
