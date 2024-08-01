package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	DbHost     string `env:"DB_HOST" envDefault:"localhost"`
	DbPort     string `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER" envDefault:"admin"`
	DbPassword string `env:"DB_PASSWORD"`
	DbName     string `env:"DB_NAME" envDefault:"database"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
