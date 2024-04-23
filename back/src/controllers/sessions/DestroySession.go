package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/gofiber/fiber/v2"
)

func DestroySession(ctx *fiber.Ctx) error {
	ctx.ClearCookie("session")

	return helpers.SendOK(ctx)
}
