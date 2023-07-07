package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB DB
}

type DB struct {
	Driver         string `envconfig:"DB_DRIVER" default:"mysql"`
	DataSourceName string `envconfig:"DB_DATA_SOURCE_NAME" default:"test:test@tcp(mysql:3306)/app?charset=utf8mb4"`
}

func New() *Config {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
