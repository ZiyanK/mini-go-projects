package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Sup bro")
	})

	// app.Get("/name/:name", func(c *fiber.Ctx) error {
	// 	msg := fmt.Sprintf("Hello %s", c.Params("name"))
	// 	return c.SendString(msg)
	// })

	app.Get("/name/no", func(c *fiber.Ctx) error {
		return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
	})

	app.Listen(":8080")

}

func letssee() {
    fmt.Println("Hello")
}

func hello() {
    fmt.Println("Thanks fenil")
}
