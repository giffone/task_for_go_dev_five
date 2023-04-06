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

func (h *Handlers) Save(c *fiber.Ctx) error {
	date := c.Params("date")
	err := h.svc.Add(c.Context(), date)
	if err != nil {
		return c.JSON(err)
	}
	ok := struct {
		Success bool `json:"success"`
	}{
		Success: true,
	}
	return c.JSON(ok)
}

func (h *Handlers) Get(c *fiber.Ctx) error {
	date := c.Params("date")
	code := c.Params("code")
	err := h.svc.Get(c.Context(), date, code)
	if err != nil {
		return c.JSON(err)
	}
	return c.JSON(nil)
}
