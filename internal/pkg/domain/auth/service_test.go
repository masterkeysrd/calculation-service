package auth_test

import (
	"testing"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/auth"
)

func TestAuthService_SignUp(t *testing.T) {
	var tests = []struct {
		name  string
		input auth.SignUpRequest
		err   error
	}{
		{
			name: "Return no error when username and password are valid",
			input: auth.SignUpRequest{
				UserName: "testing@test.com",
				Password: "12345678",
			},
			err: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s auth.Service = auth.NewService()
			err := s.SignUp(tt.input)
			if err != tt.err {
				t.Errorf("AuthService.SignUp() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}

func TestAuthService_SignIn(t *testing.T) {
	var tests = []struct {
		name  string
		input auth.SignInRequest
		want  auth.SignInResponse
		err   error
	}{
		{
			name: "Return access token and refresh token",
			input: auth.SignInRequest{
				UserName: "testing@test.com",
				Password: "12345678",
			},
			want: auth.SignInResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			err: nil,
		},
		{
			name: "Return error when username is empty",
			input: auth.SignInRequest{
				UserName: "",
				Password: "12345678",
			},
			want: auth.SignInResponse{},
			err:  auth.ErrInvalidUserName,
		},
		{
			name: "Return error when password is empty",
			input: auth.SignInRequest{
				UserName: "testing@test.com",
				Password: "",
			},
			want: auth.SignInResponse{},
			err:  auth.ErrInvalidCredentials,
		},
		{
			name: "Return error when username is invalid",
			input: auth.SignInRequest{
				UserName: "testing",
				Password: "12345678",
			},
			want: auth.SignInResponse{},
			err:  auth.ErrInvalidUserName,
		},
		{
			name: "Return error when credentials are invalid",
			input: auth.SignInRequest{
				UserName: "invalid@credentials.com",
				Password: "12345678",
			},
			want: auth.SignInResponse{},
			err:  auth.ErrInvalidCredentials,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s auth.Service = auth.NewService()
			got, err := s.SignIn(tt.input)
			if err != tt.err {
				t.Errorf("AuthService.SignIn() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got.AccessToken != tt.want.AccessToken {
				t.Errorf("AuthService.SignIn() = %v, want %v", got, tt.want)
			}
			if got.RefreshToken != tt.want.RefreshToken {
				t.Errorf("AuthService.SignIn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_RefreshToken(t *testing.T) {
	var tests = []struct {
		name  string
		input auth.RefreshRequest
		want  auth.SignInResponse
		err   error
	}{
		{
			name: "Return access token and refresh token if refresh token is valid",
			input: auth.RefreshRequest{
				RefreshToken: "refresh_token",
			},
			want: auth.SignInResponse{
				AccessToken:  "access_token",
				RefreshToken: "refresh_token",
			},
			err: nil,
		},
		{
			name: "Return error when refresh token is invalid",
			input: auth.RefreshRequest{
				RefreshToken: "invalid_refresh_token",
			},
			want: auth.SignInResponse{},
			err:  auth.ErrInvalidRefreshToken,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s auth.Service = auth.NewService()
			got, err := s.Refresh(tt.input)
			if err != tt.err {
				t.Errorf("AuthService.RefreshToken() error = %v, wantErr %v", err, tt.err)
				return
			}
			if got.AccessToken != tt.want.AccessToken {
				t.Errorf("AuthService.RefreshToken() = %v, want %v", got, tt.want)
			}
			if got.RefreshToken != tt.want.RefreshToken {
				t.Errorf("AuthService.RefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthService_SignOut(t *testing.T) {
	var test = []struct {
		name  string
		input auth.SignOutRequest
		err   error
	}{
		{
			name: "Return no error when access token is valid",
			input: auth.SignOutRequest{
				AccessToken: "access_token",
			},
			err: nil,
		},
		{
			name: "Return error when access token is invalid",
			input: auth.SignOutRequest{
				AccessToken: "invalid_access_token",
			},
			err: auth.ErrInvalidCredentials,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			var s auth.Service = auth.NewService()
			err := s.SignOut(tt.input)
			if err != tt.err {
				t.Errorf("AuthService.SignOut() error = %v, wantErr %v", err, tt.err)
				return
			}
		})
	}
}
