package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/appsynth-org/transit/service"
	"github.com/appsynth-org/transit/utils"
	"github.com/joho/godotenv"
)

func generateEnv() {
	fmt.Println("Generating .env file...")
	file, err := os.OpenFile(".env", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	defer file.Close()

	_, err = file.WriteString("SERVICE_ACCOUNT_BASE64={REPLACE_WITH_SERVICE_ACCOUNT_BASE64}\n")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}
	_, err = file.WriteString("GOOGLE_SHEET_ID={REPLACE_WITH_GOOGLE_SHEET_ID}\n")
	if err != nil {
		log.Fatal("Error opening file:", err)
	}

	fmt.Println("Generated .env file successfully! Please replace the value with your own.")
}

func main() {
	ctx := context.Background()

	path := flag.String("file", ".env", "File path to read spreadsheet")
	isGenerateEnv := flag.Bool("generate", false, "Generate .env file")
	flag.Parse()

	if *isGenerateEnv {
		generateEnv()
		return
	}

	fmt.Println("Reading the config from file:", *path)

	err := godotenv.Load(*path)
	if err != nil {
		log.Fatalf("Unable to load env config %v", err)
	}

	groups, err := service.ReadSpreadSheet(ctx)
	if err != nil {
		log.Fatalf("Unable to read spreadsheet %v", err)
	}

	/**
	*	Generate locale files and save to
	*	- ./output/iOS
	*	|- en.strings
	*	|- th.strings
	*	- ./output/Android
	*	|- en.xml
	*	|- th.xml
	**/
	utils.GenerateLocale(groups)
}
