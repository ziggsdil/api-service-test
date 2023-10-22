package config

import "github.com/ziggsdil/api-service-test/pkg/db"

type Config struct {
	Host string `config:"APP_HOST"`
	Port string `config:"APP_PORT"`

	Postgres db.Config `config:"postgres"`
}
