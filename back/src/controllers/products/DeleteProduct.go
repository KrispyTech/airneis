package controllers

import "github.com/gofiber/fiber/v2"

func DeleteProduct(ctx *fiber.Ctx) error {
	return ctx.SendString("Delete product")
}
