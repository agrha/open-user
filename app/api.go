package main

import (
	"fmt"
	"open-user/generated/api"
	"open-user/internal/app/controller"

	_ "github.com/joho/godotenv/autoload"

	"github.com/labstack/echo/v4"
)

func serveApi() error {
	e := echo.New()

	services, config, err := Initialize()
	if err != nil {
		return fmt.Errorf("error initializing: %w", err)
	}

	controller := controller.NewController(controller.NewControllerOption{
		Services: services,
		Config:   config,
	})
	registerHandlers(e, controller)

	err = e.Start(":1323")
	if err != nil {
		return fmt.Errorf("error starting echo server: %w", err)
	}
	return nil
}

func registerHandlers(e *echo.Echo, controller *controller.Controller) {
	api.RegisterHandlers(e, controller)
}
