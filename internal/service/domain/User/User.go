package user

import (
	"todoapp/internal/entities"
	"todoapp/internal/storage"
)

type UserService struct {
	UserRepository storage.UserRepository
}

func NewUserService(usrRepo storage.UserRepository) *UserService {
	return &UserService{UserRepository: usrRepo}
}

func (u *UserService) FindByUser(user entities.User) (*entities.User, error) {
	usr, err := u.UserRepository.FindByUser(user)
	if err != nil {
		return nil, err
	}
	return usr, nil
}
