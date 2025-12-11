package repository

import (
	"Shop/models"
	"context"
	"errors"

	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewShopRepository(db *gorm.DB) ShopRepository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id int64) (*models.Shop, error) {
	var shop *models.Shop
	err := r.db.WithContext(ctx).First(&shop, id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	return shop, err
}

func (r *repository) FindByUserID(ctx context.Context, userID int64) ([]models.Shop, error) {
	var shops []models.Shop
	err := r.db.WithContext(ctx).Where("user_id = ?", userID).Find(&shops).Error

	if err != nil {
		return nil, err
	}

	return shops, nil
}

func (r *repository) Create(ctx context.Context, shop *models.Shop) error {
	return r.db.WithContext(ctx).Create(shop).Error
}

func (r *repository) Update(ctx context.Context, shop *models.Shop) error {
	return r.db.WithContext(ctx).Save(shop).Error
}

func (r *repository) Delete(ctx context.Context, shop *models.Shop) error {
	return r.db.WithContext(ctx).Delete(shop).Error
}