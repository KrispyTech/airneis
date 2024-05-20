package sessions

import (
	"net/http"

	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/KrispyTech/airneis/src/controllers/users"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/gofiber/fiber/v2"
)

func CreateSession(ctx *fiber.Ctx, vlt vault.VaultClient) error {
	input, errInput := helpers.GetAuthContext(ctx)
	if errInput != nil {
		return helpers.SendInternalError(ctx, errInput)
	}

	user, errFindUser := users.GetUserByEmail(input.Email)
	if errFindUser != nil {
		return helpers.SendUnauthorized(ctx, errFindUser)
	}

	serializedUser := serializers.SerializeUser(user)
	userToken, errToken := auth.CreateToken(serializedUser, vlt)
	if errToken != nil {
		return helpers.SendInternalError(ctx, errToken)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "session",
		Value: userToken,
	})

	return helpers.SendResult(helpers.Response{
		Message: constants.OkMessage,
		Status:  http.StatusOK,
		Data:    userToken,
	}, ctx)
}
