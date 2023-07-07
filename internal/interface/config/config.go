package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	CORS CORS
}

type CORS struct {
	AllowOrigins string `envconfig:"CORS_ALLOW_ORIGINS" default:"http://localhost:3000,http://localhost:3001"`
}

func New() *Config {
	var conf Config
	err := envconfig.Process("", &conf)
	if err != nil {
		panic(err)
	}

	return &conf
}
