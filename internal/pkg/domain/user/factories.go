package user

import (
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/security/hash"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/validator"
)

type UserFactory = func(userName, password string) (*User, error)

func NewUserFactory(validator *validator.Validator) UserFactory {
	return func(userName, password string) (*User, error) {
		passwordHash, err := hash.HashPassword(password)

		if err != nil {
			return nil, err
		}

		user := &User{
			UserName: userName,
			Password: passwordHash,
			Balance: &UserBalance{
				Amount: 50,
			},
		}

		err = validator.Validate(user)

		if err != nil {
			return nil, err
		}

		return user, nil
	}
}
