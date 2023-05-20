package user

type CreateUserRequest struct {
	UserName string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

type DeleteUserRequest struct {
	UserName string `json:"username" validate:"required,email"`
}

type FindUserResponse struct {
	ID       uint64 `json:"id"`
	UserName string `json:"username"`
}

type Service interface {
	FindByUserName(username string) (*FindUserResponse, error)
	Create(request CreateUserRequest) error
	Delete(request DeleteUserRequest) error
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

func (s userService) Create(request CreateUserRequest) error {
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

func (s userService) Delete(request DeleteUserRequest) error {
	if request.UserName == "" {
		return ErrUserNameRequired
	}

	user, err := s.repository.FindByUserName(request.UserName)

	if err != nil {
		return err
	}

	if user == nil || user.ID == 0 {
		return ErrUserNotFound
	}

	return s.repository.Delete(user)
}
