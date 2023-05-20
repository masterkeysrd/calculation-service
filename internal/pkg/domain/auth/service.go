package auth

import (
	"strconv"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/jwt"
	"go.uber.org/dig"
)

type Service interface {
	SignUp(request SignUpRequest) error
	SignIn(request SignInRequest) (*SignInResponse, error)
	Refresh(request RefreshRequest) (*SignInResponse, error)
	SignOut(request SignOutRequest) error
}

type authService struct {
	userService user.Service
	jwtService  jwt.Service
}

type AuthServiceParams struct {
	dig.In
	JWTService  jwt.Service
	UserService user.Service
}

func NewAuthService(options AuthServiceParams) Service {
	return &authService{
		jwtService:  options.JWTService,
		userService: options.UserService,
	}
}

func (s *authService) SignUp(request SignUpRequest) error {
	err := s.userService.Create(user.CreateUserRequest{
		UserName: request.UserName,
		Password: request.Password,
	})

	return err
}

func (s *authService) SignIn(request SignInRequest) (*SignInResponse, error) {
	user, err := s.userService.VerifyCredentials(request.UserName, request.Password)

	if err != nil {
		return nil, err
	}

	tokens, err := s.jwtService.GenerateTokens(strconv.FormatUint(user.ID, 10))
	if err != nil {
		return nil, err
	}

	return &SignInResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}

func (s *authService) Refresh(request RefreshRequest) (*SignInResponse, error) {
	return nil, nil
}

func (s *authService) SignOut(request SignOutRequest) error {
	return nil
}
