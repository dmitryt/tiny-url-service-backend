package app

import (
	"net/http"

	"github.com/dmitryt/tiny-url-service-backend/internal/api"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/gofiber/fiber/v2"
)

type App struct {
	config *config.Config
}

type DummyResponse struct {
	OK bool
}

func New(config *config.Config, client *http.Client) *App {
	return &App{
		config: config,
	}
}

func (p *App) Run(addr string) error {
	apiInstance := api.New()
	app := fiber.New()

	api := app.Group("/api/v1")
	apiInstance.Handle(api)

	return app.Listen(addr)
}
