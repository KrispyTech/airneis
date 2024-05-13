package cart

import (
	"encoding/json"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/pkg/config"
	model "github.com/KrispyTech/airneis/src/models"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/gofiber/fiber/v2"
)

func GetCartOfLoggedUser(ctx *fiber.Ctx) (model.Order, error) {
	var user serializers.SerializedUser
	var cart model.Order

	if err := json.Unmarshal([]byte(ctx.Get("user")), &user); err != nil {
		return model.Order{}, helpers.SendUnauthorized(ctx, err)
	}

	if err := config.Database.First(&cart, "user_id = ? AND FinalizedAt = ?", user.ID, nil).Error; err != nil {
		return model.Order{}, helpers.SendNotFound(ctx, err)
	}

	return cart, nil
}

