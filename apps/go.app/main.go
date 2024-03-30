package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"kartaca.com/mod/handler"
)

func main() {

	var app *fiber.App = fiber.New()
	app.Use(cors.New())

	app.Get("/", handler.Hello())
	app.Get("/staj", handler.Countries())

	port := os.Getenv("PORT")
	if port == "" {
		port = "5555"
	}

	log.Fatal(app.Listen(":" + port))
}
