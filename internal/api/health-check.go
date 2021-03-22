package api

import (
	"github.com/gofiber/fiber/v2"
)

type DummyResponse struct {
	OK bool
}

func healthCheckHandler(c *fiber.Ctx) error {
	return c.JSON(DummyResponse{true})
}

func (p *API) HandleHealthCheck(r fiber.Router) {
	// p.HandleLinks(r.Group("/links"))
	// p.HandleLinks(r.Group("/auth"))
	r.Get("", healthCheckHandler)
}
