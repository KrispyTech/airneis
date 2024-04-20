package controllers

import "github.com/gofiber/fiber/v2"

func DestroySession(ctx *fiber.Ctx) error {
	return ctx.SendString("Destroy session")
}
