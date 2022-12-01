package services

import (
	"github.com/gofiber/fiber/v2"
)

func PostQuestions(c *fiber.Ctx) error {
	body := c.Body()
	return c.Send(body)
}
