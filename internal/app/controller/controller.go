package controller

import (
	"open-user/internal/app/config"
	"open-user/internal/app/service"
)

type Controller struct {
	Services *service.Services
	Config   *config.Config
}

type NewControllerOption struct {
	Services *service.Services
	Config   *config.Config
}

func NewController(opts NewControllerOption) *Controller {
	return &Controller{
		Services: opts.Services,
		Config:   opts.Config,
	}
}
