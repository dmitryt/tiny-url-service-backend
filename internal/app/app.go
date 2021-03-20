package app

import (
	"net/http"

	"github.com/dmitryt/tiny-url-service-backend/internal/api"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/rs/zerolog/log"
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

	r := mux.NewRouter()

	apiInstance.Handle(r.PathPrefix("/api").Subrouter())

	log.Info().Msgf("Listening at %s", addr)

	handler := cors.New(
		cors.Options{
			AllowedMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodHead},
		}).Handler(r)

	return http.ListenAndServe(addr, handler)
}
