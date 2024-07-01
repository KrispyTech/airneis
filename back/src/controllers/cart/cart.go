package cart

import (
	"encoding/json"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
)

const (
	INVALID_USER_ERROR   string = "Invalid user"
	CART_NOT_FOUND_ERROR string = "Cart not found"
)

func GetCartOfLoggedUser(ctx *fiber.Ctx) (model.Order, error) {
	var user serializers.SerializedUser
	var cart model.Order

	if err := json.Unmarshal([]byte(ctx.Get("user")), &user); err != nil {
		return model.Order{}, helpers.SendUnauthorized(ctx, errors.New(INVALID_USER_ERROR))
	}

	if err := config.Database.First(&cart, "user_id = ? AND FinalizedAt = ?", user.ID, nil).Error; err != nil {
		return model.Order{}, helpers.SendNotFound(ctx, errors.New(CART_NOT_FOUND_ERROR))
	}

	return cart, nil
}

func GetCartOfLoggedUserIfExistsOrCreateOne(ctx *fiber.Ctx) (model.Order, error) {
	cart, err := GetCartOfLoggedUser(ctx)

	if err != nil {
		if err.Error() == INVALID_USER_ERROR {
			return model.Order{}, helpers.SendUnauthorized(ctx, err)
		}

		if err.Error() == CART_NOT_FOUND_ERROR {
			var newCart serializers.SerializedUser
			if errorUnmarshall := json.Unmarshal([]byte(ctx.Get("newCart")), &newCart); errorUnmarshall != nil {
				return model.Order{}, helpers.SendUnauthorized(ctx, errors.New(INVALID_USER_ERROR))
			}

			cart = model.Order{
				UserID: newCart.ID,
			}

			if err := config.Database.Create(&cart).Error; err != nil {
				return model.Order{}, helpers.SendInternalError(ctx, err)
			}

			return cart, nil
		}

		return model.Order{}, err
	}

	return cart, nil
}
