package helpers

import (
	"net/http"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
)

func SendResult(status int, message string, ctx *fiber.Ctx) error {
	type Response struct {
		Message string `json:"message"`
		Status  int    `json:"status"`
	}

	response := Response{
		Message: message, Status: status,
	}

	return ctx.Status(status).JSON(response)
}

func SendBadRequest(ctx *fiber.Ctx) error {
	return SendResult(
		http.StatusBadRequest,
		constants.BadRequestMessage,
		ctx)
}

func SendNotFound(ctx *fiber.Ctx) error {
	return SendResult(
		http.StatusNotFound,
		constants.NotFoundMessage,
		ctx)
}

func SendInternalError(ctx *fiber.Ctx) error {
	return SendResult(
		http.StatusInternalServerError,
		constants.InternalServerErrorMessage,
		ctx)
}

func SendUnauthorized(ctx *fiber.Ctx) error {
	return SendResult(
		http.StatusUnauthorized,
		constants.UnauthorizedMessage,
		ctx)
}

func SendOK(ctx *fiber.Ctx) error {
	return SendResult(
		http.StatusOK,
		constants.OkMessage,
		ctx)
}
