package app

import (
	"net/http"

	"github.com/dmitryt/tiny-url-service-backend/internal/api"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/dmitryt/tiny-url-service-backend/internal/middleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))
	app.Use(middleware.Auth)

	api := app.Group("/api/v1")
	apiInstance.Handle(api)

	return app.Listen(addr)
}
