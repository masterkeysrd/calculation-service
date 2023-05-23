package repositories

import (
	"errors"

	"github.com/masterkeysrd/calculation-service/internal/pkg/domain/user"
	"github.com/masterkeysrd/calculation-service/internal/pkg/infra/persistence/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

func (r *userRepository) FindByID(id uint) (*user.User, error) {
	var u models.User

	err := r.db.First(&u, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, user.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       u.ID,
		UserName: u.Username,
		Password: u.Password,
	}, nil
}

func (r *userRepository) FindByUserName(userName string) (*user.User, error) {
	var u models.User
	err := r.db.Where("username = ?", userName).First(&u).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, user.ErrUserNotFound
	}

	if err != nil {
		return nil, err
	}

	return &user.User{
		ID:       u.ID,
		UserName: u.Username,
		Password: u.Password,
	}, nil
}

func (r *userRepository) Create(user *user.User) error {
	u := &models.User{
		Username: user.UserName,
		Password: user.Password,
	}

	if err := r.db.Create(u).Error; err != nil {
		return err
	}

	user.ID = u.ID
	return nil
}

func (r *userRepository) Update(user *user.User) error {
	err := r.db.Model(&models.User{}).Where("id = ?", user.ID).Updates(models.User{
		Username: user.UserName,
		Password: user.Password,
	}).Error

	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) Delete(user *user.User) error {
	if err := r.db.Delete(&models.User{}, user.ID).Error; err != nil {
		return err
	}

	return nil
}
