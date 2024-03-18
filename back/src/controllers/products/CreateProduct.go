package controllers

import "github.com/gofiber/fiber/v2"

func CreateProduct(ctx *fiber.Ctx) error {
	return ctx.SendString("Create product")
}
