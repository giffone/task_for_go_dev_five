package app

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func newRouter() *fiber.App {
	r := fiber.New()

	r.Use("/swagger/*", swagger.HandlerDefault)

	// register endpoints
	r.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	return r
}
