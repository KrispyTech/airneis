package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func DeleteCategory(ctx *fiber.Ctx) error {
	return ctx.SendString(fmt.Sprint("Delete Category %s", ctx.Params("categoryID")))
}
