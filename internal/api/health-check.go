package api

import (
	"github.com/gofiber/fiber/v2"
)

type DummyResponse struct {
	OK  bool
	Err string
}

func healthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(DummyResponse{true, ""})
}

func (p *API) HandleHealthCheck(r fiber.Router) {
	r.Get("", healthCheckHandler)
}
