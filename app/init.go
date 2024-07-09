package main

import (
	"agrha/open-user/internal/service"
)

func Initialize() (
	services *service.Services,
	configs *config.Config,
	err error,
) {
	configs = config.MustLoadConfig()
	services, err = InitializeServices(configs, fflagConfigs)
	if err != nil {
		return
	}
	return
}
