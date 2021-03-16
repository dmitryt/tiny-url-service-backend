package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"

	"github.com/dmitryt/tiny-url-service-backend/internal/app"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/rs/zerolog/log"
)

var ErrAppFatal = errors.New("application cannot start2")

func main() {
	flag.Parse()

	cfg, err := config.Read()
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
	app := app.New(cfg, http.DefaultClient)
	err = app.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
}
