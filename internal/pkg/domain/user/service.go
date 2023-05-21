package user

import "go.uber.org/dig"

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type DeleteUserRequest struct {
	UserID uint64 `json:"username" validate:"required,email"`
}

type FindUserResponse struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
}

type Service interface {
	Get(id uint64) (*FindUserResponse, error)
	GetByUserName(username string) (*FindUserResponse, error)
	Create(request CreateUserRequest) error
	Delete(request DeleteUserRequest) error
	VerifyCredentials(username string, password string) (*FindUserResponse, error)
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

func (s service) Get(id uint64) (*FindUserResponse, error) {
	user, err := s.repository.FindByID(id)

	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return &FindUserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}

func (s service) GetByUserName(userName string) (*FindUserResponse, error) {
	if userName == "" {
		return nil, ErrUserNameRequired
	}

	user, err := s.repository.FindByUserName(userName)

	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, ErrUserNotFound
	}

	return &FindUserResponse{
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

	s.repository.Create(user)

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

func (s service) VerifyCredentials(username string, password string) (*FindUserResponse, error) {
	user, err := s.repository.FindByUserName(username)

	if err != nil {
		return nil, err
	}

	if user == nil || user.ID == 0 {
		return nil, ErrInvalidCredentials
	}

	if err := user.ComparePassword(password); err != nil {
		return nil, err
	}

	return &FindUserResponse{
		ID:       user.ID,
		UserName: user.UserName,
	}, nil
}
