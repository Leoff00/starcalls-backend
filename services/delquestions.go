package services

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/starcalls-backend/db"
	"github.com/starcalls-backend/models"
)

func DeleteQuestions(c *fiber.Ctx) error {
	q_id := c.Params("q_id")
	q := new(models.Questions)

	res := db.GetDatabase().Delete(&q, "q_id = ?", q_id)

	if res.RowsAffected == 0 {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Status(http.StatusOK).JSON(&fiber.Map{
		"message": "Successfully Deleted!",
	})
}
