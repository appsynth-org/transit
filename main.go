package main

import (
	"context"
	"log"

	"github.com/appsynth-org/transit/service"
	"github.com/joho/godotenv"
)

type Translation struct {
	Th string `json:"th"`
	En string `json:"en"`
}

type Key struct {
	Comment     string      `json:"comment"`
	AndroidKey  string      `json:"android_key"`
	IosKey      string      `json:"ios_key"`
	Translation Translation `json:"translation"`
}

type Localize struct {
	GroupName string `json:"group_name"`
	Keys      []Key  `json:"keys"`
}

func main() {
	ctx := context.Background()

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load env config %v", err)
	}

	_, err = service.ReadSpreadSheet(ctx)
	if err != nil {
		log.Fatalf("Unable to read spreadsheet %v", err)
	}

	// TODO: write translation to file
}
