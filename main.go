package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"monexa/config"
	"monexa/db"
	"monexa/server"
	"net/http"
	"os"
)

func main() {
	var (
		appPort = config.CheckAndReturn(config.Config.App.Port)

		dbConfig = db.DBConfig{
			Host:     config.CheckAndReturn(config.Config.DB.Host),
			User:     config.CheckAndReturn(config.Config.DB.User),
			Password: config.CheckAndReturn(config.Config.DB.Password),
			DBName:   config.CheckAndReturn(config.Config.DB.DBName),
			Port:     config.CheckAndReturn(config.Config.DB.Port),
			SSLMode:  config.CheckAndReturn(config.Config.DB.SSLMode),
		}
	)
	fmt.Fprint(os.Stderr, "👍 [1] Environment config variables are loaded successfully\n")

	dbClient := db.ConnectDatabase(dbConfig)
	fmt.Fprint(os.Stderr, "👍 [2] Database is connected successfully\n")

	db.Migrate(dbClient)
	fmt.Fprint(os.Stderr, "👍 [3] Migrations applied successfully\n")

	e := echo.New()
	fmt.Fprint(os.Stderr, "👍 [4] Echo HTTP client is initiated successfully\n")

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Monexa!")
	})

	fmt.Fprint(os.Stderr, "👍 [5] Starting HTTP server...\n")
	server.StartServer(e, ":"+appPort)
}
