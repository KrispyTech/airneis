package sessions

import (
	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/gofiber/fiber/v2"
)

func ValidateSession(ctx *fiber.Ctx, vlt vault.VaultClient) error {
	sessionToken, ok := helpers.GetSessionCookiesFromContext(ctx)
	if !ok {
		if err := CreateSession(ctx, vlt); err != nil {
			return helpers.SendInternalError(ctx, err)
		}

		return nil
	}

	if _, err := auth.CheckToken(sessionToken, vlt); err != nil {
		return helpers.SendUnauthorized(ctx, err)
	}

	return helpers.SendOK(ctx)
}
