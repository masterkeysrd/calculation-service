package auth

import "errors"

type AuthService struct{}

func NewAuthService() Service {
	return &AuthService{}
}

func (s *AuthService) SignUp(request SignUpRequest) error {
	return nil
}

func (s *AuthService) SignIn(request SignInRequest) (SignInResponse, error) {
	return SignInResponse{}, errors.New("not implemented")
}

func (s *AuthService) Refresh(request RefreshRequest) (SignInResponse, error) {
	return SignInResponse{}, nil
}

func (s *AuthService) SignOut(request SignOutRequest) error {
	return nil
}
