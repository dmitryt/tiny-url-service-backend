package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/dmitryt/tiny-url-service-backend/internal/app"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var ErrAppFatal = errors.New("application cannot start2")

var isDev bool

func init() {
	flag.BoolVar(&isDev, "isDev", false, "Whether we run app in dev mode")
}

func main() {
	flag.Parse()

	cfg, err := config.Read()
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
	if isDev {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}
	app := app.New(cfg, http.DefaultClient)
	err = app.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
}
