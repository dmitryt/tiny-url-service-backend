package config

import (
	"context"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Config struct {
	Host     string `yaml:"host" config:"required"`
	Port     int    `yaml:"port" config:"required"`
	LogLevel string `yaml:"logLevel"`
}

func GetDefaultConfig() *Config {
	return &Config{
		Host:     "localhost",
		Port:     8082,
		LogLevel: "debug",
	}
}

func Read() (config *Config, err error) {
	config = GetDefaultConfig()
	loader := confita.NewLoader(
		env.NewBackend(),
	)
	err = loader.Load(context.Background(), config)

	return
}
