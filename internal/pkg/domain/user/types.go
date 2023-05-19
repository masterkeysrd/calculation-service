package user

import "time"

type User struct {
	ID        uint64
	UserName  string
	Password  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
}
