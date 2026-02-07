package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port string
	Name string
}

func Load() *Config {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	name := os.Getenv("SERVICE_NAME")
	if name == "" {
		name = "unknown-service"
	}

	return &Config{
		Port: port,
		Name: name,
	}
}
