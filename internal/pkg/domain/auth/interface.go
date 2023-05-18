package auth

type Service interface {
	SignUp(request SignUpRequest) error
	SignIn(request SignInRequest) (SignInResponse, error)
	Refresh(request RefreshRequest) (SignInResponse, error)
	SignOut(request SignOutRequest) error
}
