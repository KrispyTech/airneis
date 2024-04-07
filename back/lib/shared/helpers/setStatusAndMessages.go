package helpers

import "github.com/gofiber/fiber/v2"

func SetResponseStatusAndMessages(responseStatus int, responseMessage string, ctx *fiber.Ctx) error {
	ctx.Status(responseStatus)

	return ctx.SendString(responseMessage)
}
