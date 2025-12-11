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

func NewProductRepository(db *gorm.DB) ProductRepository {
	return &repository{db: db}
}

func (r *repository) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	var product *models.Product
	err := r.db.WithContext(ctx).First(&product, id).Error

	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	
	return product, err
}

func (r *repository) Update(ctx context.Context, product *models.Product) error  {
	return r.db.WithContext(ctx).Save(product).Error
}

func (r *repository) Delete(ctx context.Context, product *models.Product) error  {
	return r.db.WithContext(ctx).Delete(product).Error
}

func (r *repository) Create(ctx context.Context, product *models.Product) error  {
	return r.db.WithContext(ctx).Create(product).Error
}