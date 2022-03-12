package config

import (
	"fmt"
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Mode     string `env:"ENVIRONMENT_MODE" env-required:"true"`
	Listen   Server
	Postgres PostgresConfig
}

type Server struct {
	Port string `env:"APP_PORT" env-default:"8080"`
	HOST string `env:"APP_HOST" env-default:"localhost"`
}

type PostgresConfig struct {
	Username string `env:"PG_USER" env-required:"true"`
	Password string `env:"PG_PASSWORD" env-required:"true"`
	Host     string `env:"PG_HOST" env-default:"localhost"`
	Port     string `env:"PG_PORT" env-default:"5432"`
	Database string `env:"PG_DATABASE" env-required:"true"`
}

var (
	instance *Config
	once     sync.Once
)

func GetConfig() *Config {
	once.Do(func() {
		instance = &Config{}

		if err := cleanenv.ReadConfig(".env", instance); err != nil {
			help, _ := cleanenv.GetDescription(instance, nil)

			fmt.Print(help)
			log.Fatal(err)
		}
	})

	return instance
}
