package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/appsynth-org/transit/config"
	"github.com/appsynth-org/transit/reader"
	"github.com/appsynth-org/transit/utils"
	"github.com/appsynth-org/transit/writer"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Unable to load env config %v", err)
	}

	e.GET("/generate", func(c echo.Context) error {
		ctx := context.Background()
		groups, err := reader.ReadGoogleSheet(config, ctx)
		if err != nil {
			log.Fatalf("Unable to read spreadsheet %v", err)
		}
		writer.GenerateLocaleFiles(groups)

		return c.String(http.StatusOK, "Hello, World!")
	})

	e.GET("/", func(c echo.Context) error {
		platform := c.QueryParam("platform")
		lang := c.QueryParam("lang")

		if !utils.HasTranslation() {
			return c.String(http.StatusOK, "No translation found, please generate first")
		}

		fileName := fmt.Sprintf("%s_%s.%s", platform, lang, utils.GetFileExtension(platform))

		c.Response().Header().Set(echo.HeaderContentDisposition, fmt.Sprintf("attachment; filename=%s", fileName))
		c.Response().Header().Set(echo.HeaderContentType, c.Request().Header.Get(echo.HeaderContentType))

		location := fmt.Sprintf("../../output/%s/%s.%s", utils.GetPlatformFolderName(platform), lang, utils.GetFileExtension(platform))

		return c.File(location)
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", config.PORT)))
}
