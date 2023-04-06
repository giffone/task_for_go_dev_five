package app

import (
	"context"
	"errors"
	"nbrates/internal/api"
	"nbrates/internal/domain"
	"time"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func newRouter(h *api.Handlers) *fiber.App {
	r := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ErrorHandler: ErrHndl,
		ReadTimeout:  60 * time.Second,
		WriteTimeout: 60 * time.Second,
		BodyLimit:    -1,
	})
	r.Use(recover.New())
	r.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))
	r.Use("/swagger/*", swagger.HandlerDefault)

	// register endpoints
	curr := r.Group("/currency")
	curr.Get("/save/:date", h.Save)
	curr.Get("/get/:date/:code?", h.Get)

	return r
}

func ErrHndl(c *fiber.Ctx, err error) error {
	resp := domain.Response{
		Code:    fiber.StatusInternalServerError,
		Success: false,
		Message: err.Error(),
	}

	var e *fiber.Error
	if errors.As(err, &e) {
		resp.Code = e.Code
	} else {
		if errors.Is(err, domain.NoRates) {
			resp.Code = fiber.StatusOK
		}
		if errors.Is(err, context.DeadlineExceeded) {
			resp.Code = fiber.StatusRequestTimeout
		}
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)

	return c.Status(resp.Code).JSON(&resp)
}
