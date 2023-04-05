package app

import (
	"fmt"
	"nbrates/internal/config"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
)

type Router struct {
	app  *fiber.App
	addr string
}

func New(conf *config.AppConf) *Router {
	r := Router{
		app:  fiber.New(),
		addr: fmt.Sprintf("%s:%s", conf.Route.Host, conf.Route.Port),
	}

	r.app.Use("/swagger/*", swagger.HandlerDefault)

	// register endpoints
	r.endpoints()

	return &r
}

func (r *Router) endpoints() {
	r.app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})
}

func (r *Router) Start() error {
	return r.app.Listen(r.addr)
}
