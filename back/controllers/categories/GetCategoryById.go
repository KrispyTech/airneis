package controllers

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func GetCategoryByID(ctx *fiber.Ctx) error {
	return ctx.SendString(fmt.Sprintf("get categrory %s", ctx.Params("categoryID")))
}
