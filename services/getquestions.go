package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/db"
	"github.com/starcalls-backend/models"
)

func GetQuestions(c *fiber.Ctx) error {
	var q []models.Questions
	db.GetDatabase().Find(&q)

	return c.Status(http.StatusOK).JSON(q)
}
