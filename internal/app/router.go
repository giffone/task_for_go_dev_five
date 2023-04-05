package app

import (
	"nbrates/internal/api"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

func newRouter(h *api.Handlers) *fiber.App {
	r := fiber.New()

	r.Use("/swagger/*", swagger.HandlerDefault)

	// register endpoints
	curr := r.Group("/currency")
	curr.Get("/save/{date}", h.Save)
	curr.Get("/{date}/{*code}", h.Get)

	return r
}
