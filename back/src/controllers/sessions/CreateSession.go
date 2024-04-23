package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	controllers "github.com/KrispyTech/airneis/src/controllers/users"
	"github.com/KrispyTech/airneis/src/serializers"
	"github.com/gofiber/fiber/v2"
)

func CreateSession(ctx *fiber.Ctx) error {
	input, errInput := helpers.GetAuthContext(ctx)
	if errInput != nil {
		return helpers.SendInternalError(ctx)
	}

	user, errFindUser := controllers.GetUserByEmail(input.Email)
	if errFindUser != nil {
		return helpers.SendUnauthorized(ctx)
	}

	serializedUser := serializers.SerializeUser(user)
	userToken, errToken := auth.CreateToken(serializedUser)
	if errToken != nil {
		return helpers.SendInternalError(ctx)
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "session",
		Value: userToken,
	})

	return helpers.SendOK(ctx)
}
