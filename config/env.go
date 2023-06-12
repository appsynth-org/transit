package config

import (
	"fmt"

	"github.com/caarlos0/env/v8"
	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT                   int    `env:"PORT"`
	SERVICE_ACCOUNT_BASE64 string `env:"SERVICE_ACCOUNT_BASE64,required"`
	GOOGLE_SHEET_ID        string `env:"GOOGLE_SHEET_ID,required"`
}

func LoadConfig() (*EnvConfig, error) {
	err := godotenv.Load("../../.env")
	if err != nil {
		return nil, err
	}
	config := EnvConfig{}

	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
		return nil, err
	}

	return &config, nil
}
