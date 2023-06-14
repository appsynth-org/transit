package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT                   int
	SERVICE_ACCOUNT_BASE64 string
	GOOGLE_SHEET_ID        string
}

func LoadConfig() (*EnvConfig, error) {
	godotenv.Load()

	port, _ := strconv.Atoi(os.Getenv("PORT"))

	config := &EnvConfig{
		PORT:                   port,
		SERVICE_ACCOUNT_BASE64: os.Getenv("SERVICE_ACCOUNT_BASE64"),
		GOOGLE_SHEET_ID:        os.Getenv("GOOGLE_SHEET_ID"),
	}

	fmt.Printf("%+v", config)

	return config, nil
}

func GenerateConfig() error {
	f, err := os.Create(".env")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer f.Close()

	f.WriteString("SERVICE_ACCOUNT_BASE64=\n")
	f.WriteString("GOOGLE_SHEET_ID=\n")

	return nil
}
