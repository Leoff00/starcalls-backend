package services

import (
	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/models"
)

func GetQuestions(c *fiber.Ctx) error {
	dt := models.Questions{Question: "foo", Q_id: 1, Answers: []string{"foo"}}
	return c.JSON(dt)
}
