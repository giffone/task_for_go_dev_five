package api

import (
	"nbrates/internal/service"

	"github.com/gofiber/fiber/v2"
)

func New(svc service.Service) *Handlers {
	return &Handlers{svc: svc}
}

type Handlers struct {
	svc service.Service
}

func (h *Handlers) Save(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}

func (h *Handlers) Get(ctx *fiber.Ctx) error {
	return ctx.JSON(nil)
}
