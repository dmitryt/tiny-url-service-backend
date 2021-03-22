package api

import (
	"github.com/gofiber/fiber/v2"
)

type API struct {
}

func New() *API {
	return &API{}
}

func (p *API) Handle(r fiber.Router) {
	p.HandleHealthCheck(r.Group("/health-check"))
	p.HandleLinks(r.Group("/links"))
	// p.HandleLinks(r.Group("/auth"))
}
