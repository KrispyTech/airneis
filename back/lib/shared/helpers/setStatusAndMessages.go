package helpers

import "github.com/gofiber/fiber/v2"

func SetStatusAndMessages(status int, message string, ctx *fiber.Ctx) error {
	ctx.Status(status)

	return ctx.SendString(message)
}
