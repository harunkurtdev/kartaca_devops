package handler

import "github.com/gofiber/fiber/v2"

func Hello() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.SendString("Merhaba Go!")
	}
}

// func hello(ctx *fiber.Ctx) error {
// 	// return ctx.SendString("Merhaba Go!")
// 	return ctx.Send([]byte("Merhaba Go!"))
// }
