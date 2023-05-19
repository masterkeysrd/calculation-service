package user

import "errors"

type Service interface {
	FindByUserName(userName string) (User, error)
	FindByID(id uint64) (User, error)
	Create(user User) error
	Delete(id uint64) error
}

type userService struct{}

func NewUserService() Service {
	return &userService{}
}

func (s *userService) FindByUserName(userName string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *userService) FindByID(id uint64) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *userService) Create(user User) error {
	return errors.New("not implemented")
}

func (s *userService) Delete(id uint64) error {
	return errors.New("not implemented")
}
