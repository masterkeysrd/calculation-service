package user

import "errors"

type UserService struct{}

func NewUserService() Service {
	return &UserService{}
}

func (s *UserService) FindByUserName(userName string) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *UserService) FindByID(id uint64) (User, error) {
	return User{}, errors.New("not implemented")
}

func (s *UserService) Create(user User) error {
	return errors.New("not implemented")
}

func (s *UserService) Delete(id uint64) error {
	return errors.New("not implemented")
}
