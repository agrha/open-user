//go:build wireinject
// +build wireinject

package main

import (
	"open-user/internal/app/config"
	"open-user/internal/app/service"
	"open-user/internal/app/service/dbtx"
	"open-user/internal/app/service/user"
	"open-user/internal/db"

	"github.com/google/wire"
)

func InitializeServices(conf *config.Config) (*service.Services, error) {
	wire.Build(
		db.NewDb,
		//db transaction
		dbtx.NewDbTxRepository,
		wire.Bind(new(dbtx.DbTxRepositoryInterface), new(*dbtx.DbTxRepository)),
		//user repository
		user.NewUserService,
		wire.Bind(new(user.UserServiceInterface), new(*user.UserService)),
		wire.Struct(new(user.UserServiceOptions), "*"),
		//user service
		user.NewUserRepository,
		wire.Bind(new(user.UserRepositoryInterface), new(*user.UserRepository)),
		// Services
		wire.Struct(new(service.NewServicesOptions), "*"),
		service.NewServices,
	)
	return nil, nil
}
