package helpers

import (
	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
)

func SendResult(status int, message string, ctx *fiber.Ctx) error {
	ctx.Status(status)

	return ctx.SendString(message)
}

func SendBadRequest(ctx *fiber.Ctx) error {
	return SendResult(
		constants.BadRequestStatus,
		constants.BadRequestMessage,
		ctx)
}

func SendNotFound(ctx *fiber.Ctx) error {
	return SendResult(
		constants.NotFoundStatus,
		constants.NotFoundMessage,
		ctx)
}

func SendInternalError(ctx *fiber.Ctx) error {
	return SendResult(
		constants.InternalServerErrorStatus,
		constants.InternalServerErrorMessage,
		ctx)
}
