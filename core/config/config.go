package config

import (
	"log"
	"os"
	"time"
)

// Config holds the configuration settings for the server.
type Config struct {
	SaveInterval time.Duration
	Filename     string
	Port         string
}

// New creates and returns a new Config instance with default values.
func New() Config {
	cfg := Config{
		SaveInterval: 10 * time.Second,
		Filename:     "sessionData.json",
		Port:         ":1378",
	}

	if saveInetrval := os.Getenv("SAVE_INTERVAL"); saveInetrval != "" {
		if d, err := time.ParseDuration(saveInetrval); err == nil {
			cfg.SaveInterval = d
		} else {
			log.Printf("Invalid SAVE_INTERVAL: %s", err)
		}
	}

	if fileName := os.Getenv("FILENAME"); fileName != "" {
		cfg.Filename = fileName
	}

	if port := os.Getenv("PORT"); port != "" {
		cfg.Port = port
	}

	return cfg
}
