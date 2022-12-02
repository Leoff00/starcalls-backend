package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/db"
	"github.com/starcalls-backend/models"
)

// ! FIX: fix the put request that returned only one change in the response
func UpdateQuestions(c *fiber.Ctx) error {
	q_id := c.Params("q_id")
	var q models.Questions

	if err := c.BodyParser(&q); err != nil {
		return c.Status(http.StatusServiceUnavailable).SendString(err.Error())
	}

	db.GetDatabase().Where("q_id = ?", q_id).Updates(&q.Answers)
	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"status":  http.StatusOK,
		"answers": &q.Answers,
	})
}
