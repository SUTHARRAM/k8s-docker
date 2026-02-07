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
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	name := os.Getenv("SERVICE_NAME")
	if name == "" {
		name = "auth-service"
	}

	return &Config{
		Port: port,
		Name: name,
	}
}
