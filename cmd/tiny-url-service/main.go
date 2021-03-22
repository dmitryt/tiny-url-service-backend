package main

import (
	"errors"
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/dmitryt/tiny-url-service-backend/internal/app"
	"github.com/dmitryt/tiny-url-service-backend/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var (
	ErrAppFatal                 = errors.New("application cannot start")
	ErrCannotCreateFixturesDir  = errors.New("cannot create fixtures directory")
	ErrCannotCreateFixturesFile = errors.New("cannot create fixtures file")
)

var fixtureFiles = []string{"links.json", "users.json"}

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
	if _, err := os.Stat(cfg.FixturesPath); os.IsNotExist(err) {
		err = os.Mkdir(cfg.FixturesPath, 0o755)
		if err != nil {
			log.Fatal().Err(err).Msgf("%s: %s", ErrCannotCreateFixturesDir, err)
		}
	}
	for _, fileName := range fixtureFiles {
		fullPath := filepath.Join(cfg.FixturesPath, fileName)
		if _, err := os.Stat(fullPath); os.IsNotExist(err) {
			_, err := os.Create(fullPath)
			if err != nil {
				log.Fatal().Err(err).Msgf("%s: %s", ErrCannotCreateFixturesFile, err)
			}
		}
	}
	app := app.New(cfg, http.DefaultClient)
	err = app.Run(fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	if err != nil {
		log.Fatal().Err(err).Msgf("%s", ErrAppFatal)
	}
}
