package controllers

import "github.com/gofiber/fiber/v2"

func GetCategories(ctx *fiber.Ctx) error {
	return ctx.SendString("getCategories")
}
