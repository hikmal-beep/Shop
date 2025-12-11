package repository

import (
	"Shop/models"
	"context"
)

type ShopRepository interface {
	FindByID(ctx context.Context, id int64) (*models.Shop, error)
	Create(ctx context.Context, shop *models.Shop) error
	Update(ctx context.Context, shop *models.Shop) error
	Delete(ctx context.Context, shop *models.Shop) error
	FindByUserID(ctx context.Context, userID int64) ([]models.Shop, error)
}