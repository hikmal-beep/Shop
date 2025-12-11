package repository

import (
	"Shop/models"
	"context"
	"errors"
	"strings"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &repository{db: db}
}

func (r *repository) Register(ctx context.Context, user *models.User) error {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		// Check for duplicate email error
		if strings.Contains(err.Error(), "uni_users_email") || strings.Contains(err.Error(), "duplicate key") {
			return errors.New("email already exists")
		}
		return err
	}
	return nil
}

func (r *repository) Login(ctx context.Context, email string, password string) (*models.User, error) {
	var user *models.User
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	if !user.CheckPassword(password) {
		return nil, gorm.ErrRecordNotFound
	}

	return user, nil
}