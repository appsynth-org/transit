package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
	"google.golang.org/api/sheets/v4"
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

	cred, err := base64.StdEncoding.DecodeString(os.Getenv("SERVICE_ACCOUNT_BASE64"))
	if err != nil {
		log.Fatalf("Unable to load env config %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.JWTConfigFromJSON(cred, "https://www.googleapis.com/auth/spreadsheets.readonly")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := config.Client(ctx)

	srv, err := sheets.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}

	spreadsheetId := os.Getenv("GOOGLE_SHEET_ID")
	readRange := "SRC"
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}

	if len(resp.Values) == 0 {
		fmt.Println("No data found.")
	} else {
		var groups []Localize
		currentGroupIndex := -1

		for _, row := range resp.Values {
			if (len(row)) > 5 {
				scope := row[0]
				keyComment := row[1]
				androidKey := row[2]
				iosKey := row[3]
				th := row[4]
				en := row[5]

				// ignore group comment
				if len(scope.(string)) > 0 && scope.(string) == "Group comment" {
					continue
				}

				// create a new group
				if (len(scope.(string))) > 0 {
					groups = append(groups, Localize{
						GroupName: scope.(string),
						Keys: []Key{
							{
								Comment:    keyComment.(string),
								AndroidKey: androidKey.(string),
								IosKey:     iosKey.(string),
								Translation: Translation{
									Th: th.(string),
									En: en.(string),
								},
							},
						},
					})

					currentGroupIndex = findGroupIndex(scope.(string), groups)
					continue
				}

				// if scope is empty, add to current group
				if (len(scope.(string))) == 0 && currentGroupIndex > -1 {
					groups[currentGroupIndex].Keys = append(groups[currentGroupIndex].Keys, Key{
						Comment:     keyComment.(string),
						AndroidKey:  androidKey.(string),
						IosKey:      iosKey.(string),
						Translation: Translation{Th: th.(string), En: en.(string)},
					})
					continue
				}

			}
		}

		fmt.Printf("Read data from sheet success, Found %v groups 📝", len(groups))

		json, err := json.MarshalIndent(groups, "", "  ")
		if err != nil {
			log.Fatalf("Unable to marshal json %v", err)
		}

		fmt.Println("Writing to output.json")
		os.WriteFile("output.json", json, 0644)
	}

}

func findGroupIndex(groupName string, groups []Localize) int {
	if (len(groupName)) == 0 {
		return -1
	}

	for index, g := range groups {
		if g.GroupName == groupName {
			return index
		}
	}
	return -1
}
