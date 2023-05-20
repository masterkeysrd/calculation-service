package auth

type SignUpRequest struct {
	UserName string `json:"username" form:"username" validate:"required,email"`
	Password string `json:"password" form:"password" validate:"required,min=8,max=32"`
}

type SignInRequest struct {
	UserName string `json:"username" form:"username" binding:"required" validate:"required"`
	Password string `json:"password" form:"password" binding:"required"`
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
