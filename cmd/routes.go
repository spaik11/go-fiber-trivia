package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/spaik11/go-fiber-trivia/handlers"
)

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)

	app.Post("/fact", handlers.CreateFact)
}
