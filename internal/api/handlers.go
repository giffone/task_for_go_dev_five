package api

import (
	"nbrates/internal/domain"
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
		return err
	}
	return c.JSON(domain.Response{
		Code: fiber.StatusOK,
		Success: true,
	})
}

func (h *Handlers) Get(c *fiber.Ctx) error {
	date := c.Params("date")
	code := c.Params("code")
	itemsDTO, err := h.svc.Get(c.Context(), date, code)
	if err != nil {
		return err
	}
	return c.JSON(itemsDTO)
}
