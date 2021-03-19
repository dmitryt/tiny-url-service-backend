package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type LogConfig struct {
	Level    string `env:"LOG_LEVEL" env-default:"info"`
	FilePath string `env:"FILE_PATH" env:"LOG_FILE"`
}

type DBConfig struct {
	Host     string `env:"DB_HOST" env-default:"localhost"`
	Port     int    `env:"DB_PORT" env-default:"28017"`
	User     string `env:"DB_USER"`
	Password string `env:"DB_PASSWORD"`
	DBName   string `env:"DB_NAME"`
	RepoType string `env:"DB_REPO_TYPE" env-default:"mongo"`
}

type Config struct {
	Host string `env:"HOST" env-default:"localhost"`
	Port int    `env:"PORT" env-default:"8082"`
	// LogConfig LogConfig
	// DBConfig DBConfig
}

func Read() (*Config, error) {
	var config Config
	err := cleanenv.ReadEnv(&config)
	if err != nil {
		fmt.Printf("HHHHHH %s", err)
	}
	return &config, err
}
