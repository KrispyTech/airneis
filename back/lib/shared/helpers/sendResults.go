package helpers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
)

func SendResult(status int, message string, data interface{}, ctx *fiber.Ctx) error {
	type Response struct {
		Message string      `json:"message"`
		Status  int         `json:"status"`
		Data    interface{} `json:"data"`
	}

	response := Response{
		Message: message, Status: status, Data: data,
	}

	return ctx.Status(status).JSON(response)
}

func SendBadRequest(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(http.StatusBadRequest, constants.BadRequestMessage, nil, ctx)
}

func SendNotFound(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(http.StatusNotFound, constants.NotFoundMessage, nil, ctx)
}

func SendInternalError(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(http.StatusInternalServerError, constants.InternalServerErrorMessage, nil, ctx)
}

func SendUnauthorized(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(http.StatusUnauthorized, constants.UnauthorizedMessage, nil, ctx)
}

func SendOK(ctx *fiber.Ctx) error {
	return SendResult(http.StatusOK, constants.OkMessage, nil, ctx)
}
