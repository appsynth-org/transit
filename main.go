package main

import (
	"context"
	"log"

	"github.com/appsynth-org/transit/service"
	"github.com/appsynth-org/transit/utils"
	"github.com/joho/godotenv"
)

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load env config %v", err)
	}

	groups, err := service.ReadSpreadSheet(ctx)
	if err != nil {
		log.Fatalf("Unable to read spreadsheet %v", err)
	}

	// TODO: write translation to file
	utils.GenerateLocale(groups)
}
