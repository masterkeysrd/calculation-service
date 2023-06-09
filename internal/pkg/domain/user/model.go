package user

import "github.com/masterkeysrd/calculation-service/internal/pkg/infra/security/hash"

type User struct {
	ID       uint
	UserName string `validate:"required,email"`
	Password string `validate:"required"`
	Balance  *UserBalance
}

type UserBalance struct {
	ID             uint
	Amount         float64
	AmountInFlight float64
}

func (u *User) ComparePassword(password string) error {
	if err := hash.ComparePassword(u.Password, password); err != nil {
		return ErrInvalidCredentials
	}

	return nil
}
