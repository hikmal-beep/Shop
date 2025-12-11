package service

import (
	"Shop/models"
	"context"
)

type CreateProductData struct {
	ShopID      int64
	Product     string
	Description string
	Quantity    int
}

type UpdateProductData struct {
	ID          int64
	ShopID      int64
	Product     string
	Description string
	Quantity    int
}

type ProductService interface {
	FindByID(ctx context.Context, id int64) (*models.Product, error)
	Create(ctx context.Context, userID int64, data CreateProductData) (*models.Product, error)
	Update(ctx context.Context, userID int64, data UpdateProductData) (*models.Product, error)
	Delete(ctx context.Context, userID int64, id int64) error
}