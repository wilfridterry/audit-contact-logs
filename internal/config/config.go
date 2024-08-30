package config

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Enviroment string

	DB MongoDB

	Port int
}

type MongoDB struct {
	URI      string
	Database string
	Username string
	Password string
}

func NewConfig() (*Config, error) {
	cf := new(Config)

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := envconfig.Process("db", &cf.DB); err != nil {
		return nil, err
	}

	return cf, nil
}
