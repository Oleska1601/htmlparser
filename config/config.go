package config

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	App      App      `json:"app"`
	Server   Server   `json:"server"`
	Postgres Postgres `json:"postgres"`
	Redis    Redis    `json:"redis"`
	Logger   Logger   `json:"logger"`
}

type App struct {
	Name    string `json:"app" env:"APP_NAME"`
	Version string `json:"version" env:"APP_VERSION"`
}

type Server struct {
	Host            string        `json:"host" env:"HOST"`
	Port            int           `json:"port" env:"PORT"`
	ShutdownTimeout time.Duration `json:"shutdown_timeout" env:"SHUTDOWN_TIMEOUT"`
}

type Postgres struct {
	MaxPoolSize int    `json:"max_pool_size" env:"MAX_POOL_SIZE"`
	PgURL       string `env:"PG_URL,required"`
}

type Redis struct {
	RedisURL      string        `env:"REDIS_URL,required"`
	RedisPassword string        `env:"REDIS_PASSWORD,required"`
	TTL           time.Duration `json:"ttl" env:"TTL"`
}

type Logger struct {
	Level string `json:"level" env:"LOGGER_LEVEL"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadConfig("./config/config.yml", cfg); err != nil {
		return nil, fmt.Errorf("cleanenv.ReadConfig: %w", err)
	}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, fmt.Errorf("cleanenv.ReadEnv: %w", err)
	}
	return cfg, nil

}
