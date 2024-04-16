package helpers

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2"
)


type inputData struct {
	Password string `json:"password"`
}

func GetPasswordFromContext(ctx *fiber.Ctx) (string, error) {
	var inputData inputData

	if err := json.Unmarshal(ctx.Body(), &inputData); err != nil {
		return "", err
	}

	return inputData.Password, nil
}
