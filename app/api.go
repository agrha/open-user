package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func serveApi() error {

	e := echo.New()

	// services, config, fflagConfig, err := Initialize()
	// if err != nil {
	// 	return fmt.Errorf("error initializing: %w", err)
	// }

	// server := server.NewServer(server.NewServerOptions{
	// 	Services:    services,
	// 	Config:      config,
	// 	FFlagConfig: fflagConfig,
	// })
	// registerHandlers(e, server)
	// registerMiddlewares(e, server)

	err := e.Start(":1323")
	if err != nil {
		return fmt.Errorf("error starting echo server: %w", err)
	}
	return nil
}
