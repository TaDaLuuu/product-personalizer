package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
	"strconv"
)

func parsePort(envVar string) uint16 {
	i, err := strconv.ParseInt(envVar, 10, 16)
	if err != nil {
		log.Fatalf("Failure when initializing app, malformed port variable: %v. Expect a valid port number, got %v", err, envVar)
	}
	return uint16(i)
}

func LoadFromEnv() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Error("Error loading .env file %v", err)
	}
	return &Config{
		Postgres: &PostgresConfig{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     parsePort(os.Getenv("POSTGRES_PORT")),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Database: os.Getenv("POSTGRES_DB"),
		},
		Port: parsePort(os.Getenv("PORT")),
	}
}
