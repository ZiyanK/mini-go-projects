package main

import (
	"jwt-auth/database"
	"jwt-auth/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {

	database.Connect()

	app := fiber.New()
	routes.Setup(app)

	app.Listen(":3001")
}
