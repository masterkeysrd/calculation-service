package auth

type SignUpRequest struct {
	UserName string `json:"username" binding:"required" validate:"email;"`
	Password string `json:"password" binding:"required" validate:"min=8"`
}

type SignInRequest struct {
	UserName string `json:"username" binding:"required" validate:"email"`
	Password string `json:"password" binding:"required"`
}

type SignInResponse struct {
	AccessToken  string ` json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type RefreshRequest struct {
	RefreshToken string `json:"refresh_token" binding:"required"`
}

type SignOutRequest struct {
	AccessToken string `header:"access_token" json:"access_tokens" binding:"required"`
}
