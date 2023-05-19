package auth

import "errors"

type Service interface {
	SignUp(request SignUpRequest) error
	SignIn(request SignInRequest) (SignInResponse, error)
	Refresh(request RefreshRequest) (SignInResponse, error)
	SignOut(request SignOutRequest) error
}

type authService struct{}

func NewAuthService() Service {
	return &authService{}
}

func (s *authService) SignUp(request SignUpRequest) error {
	return nil
}

func (s *authService) SignIn(request SignInRequest) (SignInResponse, error) {
	return SignInResponse{}, errors.New("not implemented")
}

func (s *authService) Refresh(request RefreshRequest) (SignInResponse, error) {
	return SignInResponse{}, nil
}

func (s *authService) SignOut(request SignOutRequest) error {
	return nil
}
