package user

import (
	"errors"
)

type Repository interface {
	FindByID(id uint64) (*User, error)
	FindByUserName(userName string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(user *User) error
}

type FakeUserRepository struct {
	counter uint64
	users   []*User
}

func NewFakeUserRepository() Repository {
	return &FakeUserRepository{
		counter: 1,
		users: []*User{
			{
				ID:       1,
				UserName: "admin@test.com",
				Password: "admin",
			},
		},
	}
}

func (r *FakeUserRepository) FindByID(id uint64) (*User, error) {
	for _, u := range r.users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, nil
}

func (r *FakeUserRepository) FindByUserName(userName string) (*User, error) {
	for _, u := range r.users {
		if u.UserName == userName {
			return u, nil
		}
	}

	return nil, nil
}

func (r *FakeUserRepository) Create(user *User) error {
	r.counter++
	user.ID = r.counter
	r.users = append(r.users, user)

	return nil
}

func (r *FakeUserRepository) Update(user *User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users[i] = user
		}
	}

	return nil
}

func (r *FakeUserRepository) Delete(user *User) error {
	for i, u := range r.users {
		if u.ID == user.ID {
			r.users = append(r.users[:i], r.users[i+1:]...)
			return nil
		}
	}

	return errors.New("user not found")
}
