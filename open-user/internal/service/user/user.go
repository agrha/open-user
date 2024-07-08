package user

type UserService interface {
	HelloUser() error
}

type userServiceImpl struct{}

func (s *userServiceImpl) HelloUser() error {
	// Implementation
	return nil
}

func NewService1() UserService {
	return &userServiceImpl{}
}
