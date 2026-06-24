package router

import (
	"go-mailer/letters/app/http/public_api/handler"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.GetIndex)
	app.Post("/users", handler.PostUser)
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
}
