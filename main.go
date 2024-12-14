package main

import (
	"github.com/labstack/echo/v4"
	"monexa/config"
	"monexa/migrations"
	"net/http"
)

func main() {
	e := echo.New()

	config.ConnectDatabase()

	migrations.Migrate(config.DB)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Monexa!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
