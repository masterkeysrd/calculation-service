package user

import (
	"errors"
	"fmt"
)

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type FindUserResponse struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
}

type Service interface {
	FindByUserName(username string) (*FindUserResponse, error)
	Create(user CreateUserRequest) error
	Delete(username string) error
}

type userService struct {
	repository        Repository
	createUserFactory UserFactory
}

type UserServiceOptions struct {
	Repository        Repository
	CreateUserFactory UserFactory
}

func NewUserService(options UserServiceOptions) Service {
	return &userService{
		createUserFactory: options.CreateUserFactory,
		repository:        options.Repository,
	}
}

func (s userService) FindByUserName(userName string) (*FindUserResponse, error) {
	user, err := s.repository.FindByUserName(userName)

	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == 0 {
		// TODO: Create a custom error
		return nil, errors.New("user not found")
	}

	return &FindUserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}

func (s userService) Create(request CreateUserRequest) error {
	fmt.Println("userService.Create", request)
	user, err := s.createUserFactory(request.UserName, request.Password)

	if err != nil {
		return err
	}

	s.repository.Create(user)

	return nil
}

func (s userService) Delete(username string) error {
	if username == "" {
		return errors.New("username is required")
	}

	user, err := s.repository.FindByUserName(username)

	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 {
		return errors.New("user not found")
	}

	return s.repository.Delete(user)
}
