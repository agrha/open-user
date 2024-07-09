package main

import (
	"agrha/open-user/internal/app/config"
	"agrha/open-user/internal/service"

	"github.com/google/wire"
)

func InitializeServices(conf *config.Config) (*service.Services, error) {
	wire.Build()
	return nil, nil
}
