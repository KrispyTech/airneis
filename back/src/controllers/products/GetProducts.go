package controllers

import "github.com/gofiber/fiber/v2"

func GetProducts(ctx *fiber.Ctx) error {
	return ctx.SendString("Get products")
}
