package app

import (
	"errors"
	"net/http"

	"github.com/dmitryt/tiny-url-service-backend/internal/api"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/rs/zerolog/log"
)

var (
	ErrImageFetch         = errors.New("error during fetching the image")
	ErrImageResize        = errors.New("error during resizing the image")
	ErrInvalidURI         = errors.New("invalid URI. Expected format is: /<method>/<width>/<height>/<external url>")
	ErrImageCopyFromCache = errors.New("error during copying the image from cache")
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
	mux := http.NewServeMux()
	apiInstance := api.New()

	mux.HandleFunc("/health-check", apiInstance.HealthCheckHandler)

	log.Info().Msgf("Listening at %s", addr)

	return http.ListenAndServe(addr, mux)
}
