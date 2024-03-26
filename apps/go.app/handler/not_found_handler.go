package handler

import "github.com/gofiber/fiber/v2"

func NotFound() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"message": "Not found",
		})
	}
}
