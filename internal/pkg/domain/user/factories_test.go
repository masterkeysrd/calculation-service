package user_test

import (
	"reflect"
	"testing"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
)

func TestNewUserFactory(t *testing.T) {
	var tests = []struct {
		name   string
		input  user.CreateUserRequest
		errors validator.ErrorsMap
	}{
		{
			name: "Return error when username is empty",
			input: user.CreateUserRequest{
				UserName: "",
				Password: "12345678",
			},
			errors: validator.ErrorsMap{
				"UserName": []string{"UserName field is required"},
			},
		},
		{
			name: "Return error when password is empty",
			input: user.CreateUserRequest{
				UserName: "testing@test.com",
				Password: "",
			},
			errors: validator.ErrorsMap{
				"Password": []string{"Password field is required"},
			},
		},
		{
			name: "Return error when username and password are empty",
			input: user.CreateUserRequest{
				UserName: "",
				Password: "",
			},
			errors: validator.ErrorsMap{
				"UserName": []string{"UserName field is required"},
				"Password": []string{"Password field is required"},
			},
		},
		{
			name: "Return error when username is invalid",
			input: user.CreateUserRequest{
				UserName: "testing",
				Password: "12345678",
			},
			errors: validator.ErrorsMap{
				"UserName": []string{"UserName field is not a valid email address"},
			},
		},
		{
			name: "Return error when password is shorter than 8 characters",
			input: user.CreateUserRequest{
				UserName: "testing@test.com",
				Password: "12",
			},
			errors: validator.ErrorsMap{
				"Password": []string{"Password must be at least 8 characters in length"},
			},
		},
		{
			name: "Return error when password is longer than 32 characters",
			input: user.CreateUserRequest{
				UserName: "testing@test.com",
				Password: "123456789012345678901234567890123",
			},
			errors: validator.ErrorsMap{
				"Password": []string{"Password must be a maximum of 32 characters in length"},
			},
		},
		{
			name: "Return no error when username and password are valid",
			input: user.CreateUserRequest{
				UserName: "testing@test.com",
				Password: "12345678",
			},
			errors: nil,
		},
	}

	v := validator.NewValidator()
	v.RegisterDefaultTranslations()
	createUser := user.NewUserFactory(v)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			user, err := createUser(tt.input.UserName, tt.input.Password)

			if err == nil {
				return
			}

			if v, ok := err.(*validator.ValidationErrors); ok {
				if !reflect.DeepEqual(v.Errors(), tt.errors) {
					t.Errorf("UserFactory() = %v, want %v", v.Errors(), tt.errors)
				}
				return
			}

			if err != nil {
				t.Errorf("UserFactory() error = %v", err)
			}

			if user == nil {
				t.Errorf("UserFactory() = %v, want %v", user, tt.input)
			}

			if user.UserName != tt.input.UserName {
				t.Errorf("UserFactory() = %v, want %v", user, tt.input)
			}

			// TODO: Encrypt password
			if user.Password != tt.input.Password {
				t.Errorf("UserFactory() = %v, want %v", user, tt.input)
			}
		})
	}
}
