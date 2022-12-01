package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/starcalls-backend/models"
)

func GetQuestions(c *fiber.Ctx) error {
	dt := models.Questions{Question: "foo", Q_id: uuid.UUID{}, Answers: []string{"foo"}}
	return c.JSON(dt)
}
