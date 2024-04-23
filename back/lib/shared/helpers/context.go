package helpers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)

type AuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func GetAuthContext(ctx *fiber.Ctx) (AuthData, error) {
	var authData AuthData

	if err := json.Unmarshal(ctx.Body(), &authData); err != nil {
		return AuthData{}, err
	}

	return authData, nil
}
