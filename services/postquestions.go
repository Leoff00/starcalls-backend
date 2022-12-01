package services

import (
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
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

	errCreate := db.GetDatabase().Create(&q).Error

	if errCreate != nil {
		c.Status(http.StatusBadRequest).JSON(
			&fiber.Map{"message": "could not create book"})
		return errCreate
	}

	defer db.CloseDB(db.GetDatabase())
	return c.Status(http.StatusCreated).JSON(q)

}
