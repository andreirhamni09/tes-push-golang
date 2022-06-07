package config

import (
	"os"
)

const (
	LOCAL = "local"
)

const ENVIRONMENT string = LOCAL

var env = map[string]map[string]string{
	// local environment configuration
	"local": {
		"MYSQL_PORT": "3306",
		"MYSQL_HOST": "http://127.0.0.1/",
		"MYSQL_DB":   "andre",
		"MYSQL_USER": "andre",
		"MYSQL_PASS": "0s7cosUcjNWmOHbF",
	},
}

var CONFIG = env[LOCAL]

func Getenv(key string, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

func InitConfig() {
	for key := range CONFIG {
		CONFIG[key] = Getenv(key, CONFIG[key])
		os.Setenv(key, CONFIG[key])
	}
}
