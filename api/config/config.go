package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
	"sync"
)

type (
	// Config represents configuration for the application.
	Config struct {
		App   App
		Redis Redis
	}

	// App represents general app configuration.
	App struct {
		BaseURL string `env:"APP_BASE_URL"    env-default:"localhost:8080"`
		URLLen  int    `env:"URL_LEN"    env-default:"5"`
	}

	// Redis represents redis configuration.
	Redis struct {
		Address  string `env:"R_ADDRESS" env-default:"localhost:6379"`
		Password string `env:"R_PASSWORD" env-default:""`
		DB       int    `env:"R_DB" env-default:"0"`
	}
)

var (
	config Config
	once   sync.Once
)

// Get returns config.
func Get() *Config {
	once.Do(func() {
		err := cleanenv.ReadEnv(&config)
		if err != nil {
			log.Fatal("failed to read env", err)
		}
	})

	return &config
}
