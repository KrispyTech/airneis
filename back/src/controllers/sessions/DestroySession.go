package controllers

import (
	"github.com/KrispyTech/airneis/lib/shared/helpers"
	"github.com/gofiber/fiber/v2"
)

func DestroySession(ctx *fiber.Ctx) error {
	ctx.Cookie(&fiber.Cookie{
		Name:  "cookie",
		Value: "",
	})

	return helpers.SendOk(ctx)
}
