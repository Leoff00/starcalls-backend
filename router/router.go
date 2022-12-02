package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/services"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", services.GetQuestions)
	app.Post("/", services.PostQuestions)
	app.Put("/:q_id", services.UpdateQuestions)
	app.Delete("/:q_id", services.DeleteQuestions)
}
