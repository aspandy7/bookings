package config

import (
	"errors"
	"os"
)

type Config struct {
	DatabaseURL string
	ServerPort  string
}

func Load() (Config, error) {
	url := os.Getenv("DB_URL_LOCAL")
	if url == "" {
		return Config{}, errors.New("DB_URL_LOCAL is not set")
	}

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8080"
	}

	return Config{DatabaseURL: url, ServerPort: port}, nil
}
