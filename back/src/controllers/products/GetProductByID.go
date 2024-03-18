package controllers

import "github.com/gofiber/fiber/v2"

func GetProductById(ctx *fiber.Ctx) error {
	return ctx.SendString("get product by ID")
}
