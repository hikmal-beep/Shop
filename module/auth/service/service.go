package service

import (
	"Shop/models"
	"context"
)

type RegisterUserData struct {
	Name     string
	Email    string
	Password string
}

type LoginUserData struct {
	Email    string
	Password string
}

type UserService interface {
	Register(ctx context.Context, data RegisterUserData) (*models.User, error)
	Login(ctx context.Context, data LoginUserData) (*models.User, error)
}