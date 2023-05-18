package auth

type AuthService struct{}

func NewAuthService() Service {
	return &AuthService{}
}

func (s *AuthService) SignUp(request SignUpRequest) error {
	return nil
}

func (s *AuthService) SignIn(request SignInRequest) (SignInResponse, error) {
	return SignInResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil
}

func (s *AuthService) Refresh(request RefreshRequest) (SignInResponse, error) {
	return SignInResponse{
		AccessToken:  "access_token",
		RefreshToken: "refresh_token",
	}, nil
}

func (s *AuthService) SignOut(request SignOutRequest) error {
	return nil
}
