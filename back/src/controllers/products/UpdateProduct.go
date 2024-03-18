package controllers

import "github.com/gofiber/fiber/v2"

func UpdateProduct(ctx *fiber.Ctx) error {

	return ctx.SendString("update product")

}
