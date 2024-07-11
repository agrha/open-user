package main

import (
	"log"
	"open-user/internal/app/config"
	"open-user/internal/app/service"

	"github.com/joho/godotenv"
)

func Initialize() (
	services *service.Services,
	configs *config.Config,
	err error,
) {
	// Load the .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}
	configs = config.MustLoadConfig()
	services, err = InitializeServices(configs)
	if err != nil {
		return
	}
	return
}
