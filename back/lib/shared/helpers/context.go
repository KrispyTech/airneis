package helpers

import (
	"encoding/json"

	"github.com/KrispyTech/airneis/src/serializers"

	log "github.com/sirupsen/logrus"

	"github.com/gofiber/fiber/v2"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAuthContext(ctx *fiber.Ctx) (AuthData, error) {
	log.Info(string(ctx.Body()))

	var authData AuthData
	if err := json.Unmarshal(ctx.Body(), &authData); err != nil {
		return AuthData{}, err
	}

	return authData, nil
}

func GetSessionCookiesFromContext(ctx *fiber.Ctx) (string, bool) {
	sessionCookie := ctx.Cookies("session")
	if sessionCookie == "" {
		log.Warn("Session cookie not found.")

		return "", false
	}

	return sessionCookie, true
}

func AddUserToContext(ctx *fiber.Ctx, user serializers.SerializedUser) error {
	userData, err := json.Marshal(user)
	if err != nil {
		return SendInternalError(ctx, err)
	}

	ctx.Locals("user", userData)

	return nil
}
