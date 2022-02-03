package routes

import (
	"jwt-auth/controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/register", controllers.Register)
	app.Post("/login", controllers.Login)
	app.Get("/user", controllers.GetUser)
	app.Get("/logout", controllers.Logout)
}
