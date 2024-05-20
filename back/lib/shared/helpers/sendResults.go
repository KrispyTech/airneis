package helpers

import (
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/KrispyTech/airneis/lib/shared/constants"
	"github.com/gofiber/fiber/v2"
)

type Response struct {
	Message string      `json:"message"`
	Status  int         `json:"status"`
	Data    interface{} `json:"data"`
}

func SendResult(response Response, ctx *fiber.Ctx) error {
	return ctx.Status(response.Status).JSON(response)
}

func SendBadRequest(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(Response{
		Data:    nil,
		Message: constants.BadRequestMessage,
		Status:  http.StatusBadRequest,
	}, ctx)
}

func SendNotFound(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(Response{
		Data:    nil,
		Message: constants.NotFoundMessage,
		Status:  http.StatusNotFound,
	}, ctx)
}

func SendInternalError(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(Response{
		Data:    nil,
		Message: constants.InternalServerErrorMessage,
		Status:  http.StatusInternalServerError,
	}, ctx)
}

func SendUnauthorized(ctx *fiber.Ctx, err error) error {
	log.Error(err)

	return SendResult(Response{
		Data:    nil,
		Message: constants.UnauthorizedMessage,
		Status:  http.StatusUnauthorized,
	}, ctx)
}

func SendOK(ctx *fiber.Ctx) error {
	return SendResult(Response{
		Data:    nil,
		Message: constants.OkMessage,
		Status:  http.StatusOK,
	}, ctx)
}
