package user

type User struct {
	ID       uint64
	UserName string `validate:"required,email"`
	Password string `validate:"required,min=8,max=32"`
}

func (u *User) ComparePassword(password string) error {
	if u.Password != password {
		return ErrInvalidCredentials
	}

	return nil
}
