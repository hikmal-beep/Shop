package repository

import (
	"Shop/models"
	"context"
)

type UserRepository interface {
	Register(ctx context.Context, user *models.User) error
	Login(ctx context.Context, email string, password string) (*models.User, error)
}