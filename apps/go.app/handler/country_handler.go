package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"kartaca.com/mod/repository"
)

func Countries() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		randomCity, err := repository.GetRandomCity()
		if err != nil {
			return ctx.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return ctx.JSON(randomCity)
	}
}
