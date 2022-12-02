package services

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/starcalls-backend/db"
	"github.com/starcalls-backend/models"
)

func PostQuestions(c *fiber.Ctx) error {
	q := new(models.Questions)

	err := c.BodyParser(q)

	if err != nil {
		log.Fatal("Request failed.", err.Error())
		return c.SendStatus(http.StatusBadRequest)
	}

	q.Q_ID = uuid.New()
	errCreate := db.GetDatabase().Create(&q).Error

	if errCreate != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "Could not create a new question"})
		return errCreate
	}

	return c.Status(http.StatusCreated).JSON(q)
}
