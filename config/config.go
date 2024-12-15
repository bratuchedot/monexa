package config

import (
	"fmt"
	"github.com/jinzhu/configor"
	"os"
)

var Config = struct {
	App struct {
		Port string
	}

	DB struct {
		Host     string
		User     string
		Password string
		DBName   string
		Port     string
		SSLMode  string
	}

	JWT struct {
		Secret string
	}
}{}

func init() {
	switch configor.ENV() {
	case "development":
		configor.New(&configor.Config{Environment: "development"}).Load(&Config, "config/config_development.yaml")
	}
}

func CheckAndReturn(value string) string {
	if value == "" {
		fmt.Fprint(os.Stderr, "⛔ ️Exit!!! Environment config variable missing\n")
		os.Exit(1)
	}
	//fmt.Fprint(os.Stderr, "✅ Check and return environment config variable: ", value, "\n")
	return value
}
