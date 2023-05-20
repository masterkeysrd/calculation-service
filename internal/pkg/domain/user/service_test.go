package user_test

import (
	"errors"
	"testing"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/stretchr/testify/assert"
)

type UserRepositoryStub struct {
	FindByIDFunc       func(id uint64) (*user.User, error)
	FindByUserNameFunc func(userName string) (*user.User, error)
	CreateFunc         func(user *user.User) error
	UpdateFunc         func(user *user.User) error
	DeleteFunc         func(user *user.User) error
}

func (r *UserRepositoryStub) FindByID(id uint64) (*user.User, error) {
	return r.FindByIDFunc(id)
}

func (r *UserRepositoryStub) FindByUserName(userName string) (*user.User, error) {
	return r.FindByUserNameFunc(userName)
}

func (r *UserRepositoryStub) Create(user *user.User) error {
	return r.CreateFunc(user)
}

func (r *UserRepositoryStub) Update(user *user.User) error {
	return r.UpdateFunc(user)
}

func (r *UserRepositoryStub) Delete(user *user.User) error {
	return r.DeleteFunc(user)
}

func createUserFactoryStub(err error) user.UserFactory {
	return func(userName, password string) (*user.User, error) {
		if err != nil {
			return nil, err
		}

		return &user.User{
			UserName: userName,
			Password: password,
		}, nil
	}
}

func TestUserService_FindByUserName(t *testing.T) {
	var tests = []struct {
		name       string
		input      string
		err        error
		repository user.Repository
	}{
		{
			name:  "should return a error when username is empty",
			input: "",
			err:   user.ErrUserNameRequired,
		},
		{
			name:  "should return a error when repository returns error",
			input: "username",
			err:   errors.New("repository error"),
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, errors.New("repository error")
				},
			},
		},
		{
			name:  "should return a error when user is not found",
			input: "username",
			err:   user.ErrUserNotFound,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, nil
				},
			},
		},
		{
			name:  "should return a user when user is found",
			input: "username",
			err:   nil,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return &user.User{
						ID:       1,
						UserName: userName,
					}, nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s user.Service = user.NewUserService(user.UserServiceOptions{
				Repository: tt.repository,
			})

			user, err := s.FindByUserName(tt.input)

			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.NoError(t, err)
			assert.NotNil(t, user)
			assert.Equal(t, tt.input, user.UserName)
		})
	}
}

func TestUserService_Create(t *testing.T) {
	var tests = []struct {
		name       string
		input      user.CreateUserRequest
		err        error
		factoryErr error
		repository user.Repository
	}{
		{
			name:       "should return error when factory returns error",
			input:      user.CreateUserRequest{},
			err:        errors.New("factory error"),
			factoryErr: errors.New("factory error"),
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, nil
				},
			},
		},
		{
			name:       "should return error when repository returns error",
			input:      user.CreateUserRequest{},
			err:        errors.New("repository error"),
			factoryErr: nil,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, nil
				},
				CreateFunc: func(user *user.User) error {
					return errors.New("repository error")
				},
			},
		},
		{
			name:       "should return error when user already exists",
			input:      user.CreateUserRequest{},
			err:        user.ErrUserAlreadyExists,
			factoryErr: nil,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return &user.User{}, nil
				},
			},
		},
		{
			name:       "should return no error when user is created successfully",
			input:      user.CreateUserRequest{},
			err:        nil,
			factoryErr: nil,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, nil
				},
				CreateFunc: func(user *user.User) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s user.Service = user.NewUserService(user.UserServiceOptions{
				Repository:        tt.repository,
				CreateUserFactory: createUserFactoryStub(tt.factoryErr),
			})

			err := s.Create(tt.input)

			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}

func TestUserService_Delete(t *testing.T) {
	var tests = []struct {
		name       string
		input      user.DeleteUserRequest
		err        error
		repository user.Repository
	}{
		{
			name: "should return error when username is empty",
			input: user.DeleteUserRequest{
				UserName: "",
			},
			err: user.ErrUserNameRequired,
		},
		{
			name: "should return error when repository.FindByUserName returns error",
			input: user.DeleteUserRequest{
				UserName: "username",
			},
			err: errors.New("repository error"),
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, errors.New("repository error")
				},
			},
		},
		{
			name: "should return error when user is not found",
			input: user.DeleteUserRequest{
				UserName: "username",
			},
			err: user.ErrUserNotFound,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return nil, nil
				},
			},
		},
		{
			name: "should return error when repository.Delete returns error",
			input: user.DeleteUserRequest{
				UserName: "username",
			},
			err: errors.New("repository error"),
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return &user.User{
						ID:       1,
						UserName: userName,
					}, nil
				},
				DeleteFunc: func(user *user.User) error {
					return errors.New("repository error")
				},
			},
		},
		{
			name: "should return no error when user is deleted successfully",
			input: user.DeleteUserRequest{
				UserName: "username",
			},
			err: nil,
			repository: &UserRepositoryStub{
				FindByUserNameFunc: func(userName string) (*user.User, error) {
					return &user.User{
						ID:       1,
						UserName: userName,
					}, nil
				},
				DeleteFunc: func(user *user.User) error {
					return nil
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var s user.Service = user.NewUserService(user.UserServiceOptions{
				Repository: tt.repository,
			})

			err := s.Delete(tt.input)

			if err != nil {
				assert.EqualError(t, err, tt.err.Error())
				return
			}

			assert.NoError(t, err)
		})
	}
}
