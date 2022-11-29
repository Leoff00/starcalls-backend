package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	router "github.com/starcalls-backend/router"
)

const PORT = ":3001"

func main() {
	app := fiber.New()
	app.Use(cors.New())

	router.SetupRoutes(app)

	app.Listen(":3000")
}
