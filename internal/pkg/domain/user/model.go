package user

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/security/hash"

type User struct {
	ID       uint64
	UserName string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

func (u *User) ComparePassword(password string) error {
	if err := hash.ComparePassword(u.Password, password); err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
