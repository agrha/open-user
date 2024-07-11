package user

type UserService struct {
	UserRepository UserRepositoryInterface
}

type UserServiceOptions struct {
	UserRepository UserRepositoryInterface
}

func NewUserService(options UserServiceOptions) *UserService {
	return &UserService{
		UserRepository: options.UserRepository,
	}
}
