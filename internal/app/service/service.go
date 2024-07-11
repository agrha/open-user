package service

import "open-user/internal/app/service/user"

type Services struct {
	UserService user.UserServiceInterface
}

type NewServicesOptions struct {
	UserService user.UserServiceInterface
}

func NewServices(opts NewServicesOptions) (*Services, error) {
	return &Services{
		UserService: opts.UserService,
	}, nil
}
