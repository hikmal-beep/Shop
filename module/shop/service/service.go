package service

import (
	"Shop/models"
	"context"
)

type CreateShopData struct {
	Name    string
	Address string
}

type UpdateShopData struct {
	ID      int64
	Name    string
	Address string
}

type ShopService interface {
	FindByID(ctx context.Context, id int64) (*models.Shop, error)
	FindByUserID(ctx context.Context, userID int64) (*models.Shop, error)
	Create(ctx context.Context, userID int64, data CreateShopData) (*models.Shop, error)
	Update(ctx context.Context, userID int64, data UpdateShopData) (*models.Shop, error)
	Delete(ctx context.Context, userID int64, id int64) error
}