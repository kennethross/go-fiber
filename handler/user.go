package handler

import "github.com/gofiber/fiber/v2"

func GetUsers(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": fiber.StatusOK, "message": "Get Users successful", "data": ""})
}
