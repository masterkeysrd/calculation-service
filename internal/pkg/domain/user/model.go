package user

import (
	"time"
)

type User struct {
	ID        uint64
	UserName  string `validate:"required,email"`
	Password  string `validate:"required,min=8,max=32"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
