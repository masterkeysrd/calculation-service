package user

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
)

type UserFactory = func(userName, password string) (*User, error)

func NewUserFactory(validator *validator.Validator) UserFactory {
	return func(userName, password string) (*User, error) {
		user := &User{
			UserName: userName,
			Password: password,
		}

		err := validator.Validate(user)

		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
