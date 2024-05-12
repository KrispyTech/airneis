package lib

import (
	"encoding/json"
	"fmt"

	"github.com/KrispyTech/airneis/lib/shared/auth"
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/KrispyTech/airneis/lib/shared/vault"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func (airneis *airneisFiberApp) setupLogger() fiber.Router {
	return airneis.app.Use(logger.New())
}

func (airneis *airneisFiberApp) setupCORS() fiber.Router {
	return airneis.app.Use(cors.New())
}

func authMiddleware(vc vault.VaultClient) func(*fiber.Ctx) error {
	return func(ctx *fiber.Ctx) error {
		sessionToken, ok := helpers.GetSessionCookiesFromContext(ctx)
		if !ok {
			return helpers.SendUnauthorized(ctx, fmt.Errorf("unable to get session token"))
		}

		user, errToken := auth.SendSerializedUser(sessionToken, vc)
		if errToken != nil {
			return helpers.SendUnauthorized(ctx, errToken)
		}

		if !user.IsActive {
			return helpers.SendUnauthorized(ctx, fmt.Errorf("user is inactive"))
		}

		userData, err := json.Marshal(user)
		if err != nil {
			return helpers.SendInternalError(ctx, err)
		}

		ctx.Locals("user", userData)

		return ctx.Next()
	}
}
