package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/services"
)

func SetupRoutes(app *fiber.App) {

	app.Get("/", services.GetQuestions)
}
