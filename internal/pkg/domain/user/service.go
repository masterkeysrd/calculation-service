package user

import (
	"errors"
	"fmt"

	"go.uber.org/dig"
)

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type DeleteUserRequest struct {
	UserID uint `json:"username" validate:"required,email"`
}

type UserResponse struct {
	ID       uint                `json:"id"`
	UserName string              `json:"username"`
	Balance  UserBalanceResponse `json:"balance"`
}

type UserBalanceResponse struct {
	Amount         float64 `json:"amount"`
	AmountInFlight float64 `json:"amountInFlight"`
}

type Service interface {
	Get(id uint) (*UserResponse, error)
	GetByUserName(username string) (*UserResponse, error)
	Create(request CreateUserRequest) error
	Delete(request DeleteUserRequest) error
	GetBalance(userId uint) (*UserBalanceResponse, error)
	VerifyCredentials(username string, password string) (*UserResponse, error)
}

type service struct {
	repository        Repository
	createUserFactory UserFactory
}

type UserServiceParams struct {
	dig.In
	Repository        Repository
	CreateUserFactory UserFactory
}

func NewService(options UserServiceParams) Service {
	return &service{
		createUserFactory: options.CreateUserFactory,
		repository:        options.Repository,
	}
}

func (s service) Get(id uint) (*UserResponse, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	fmt.Println("Get user.UserName=", user.UserName)

	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
		Balance: UserBalanceResponse{
			Amount:         user.Balance.Amount,
			AmountInFlight: user.Balance.AmountInFlight,
		},
	}, nil
}

func (s service) GetByUserName(userName string) (*UserResponse, error) {
	if userName == "" {
		return nil, ErrUserNameRequired
	}

	user, err := s.repository.FindByUserName(userName)

	if err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}

func (s service) Create(request CreateUserRequest) error {
	if user, _ := s.repository.FindByUserName(request.UserName); user != nil {
		return ErrUserAlreadyExists
	}

	user, err := s.createUserFactory(request.UserName, request.Password)
	if err != nil {
		return err
	}

	err = s.repository.Create(user)
	if err != nil {
		return err
	}

	return nil
}

func (s service) Delete(request DeleteUserRequest) error {
	if request.UserID == 0 {
		return ErrUserNameRequired
	}

	user, err := s.repository.FindByID(request.UserID)

	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 {
		return ErrUserNotFound
	}

	return s.repository.Delete(user)
}

func (s service) VerifyCredentials(username string, password string) (*UserResponse, error) {
	user, err := s.repository.FindByUserName(username)

	if errors.Is(err, ErrUserNotFound) {
		return nil, ErrInvalidCredentials
	}

	if err != nil {
		return nil, err
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, err
	}

	return &UserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}

func (s service) GetBalance(userId uint) (*UserBalanceResponse, error) {
	user, err := s.repository.FindByID(userId)

	if err != nil {
		return nil, err
	}

	if user.Balance == nil {
		return nil, ErrUserBalanceNotFound
	}

	return &UserBalanceResponse{
		Amount:         user.Balance.Amount,
		AmountInFlight: user.Balance.AmountInFlight,
	}, nil
}
