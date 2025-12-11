package service

import (
	"Shop/models"
	"Shop/module/product/repository"
	shopRepository "Shop/module/shop/repository"
	"context"
	"errors"
)

type service struct {
	repo     repository.ProductRepository
	shopRepo shopRepository.ShopRepository
}

func NewProductService(repo repository.ProductRepository, shopRepo shopRepository.ShopRepository) ProductService {
	return &service{
		repo:     repo,
		shopRepo: shopRepo,
	}
}

func (s *service) FindByID(ctx context.Context, id int64) (*models.Product, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) Create(ctx context.Context, userID int64, data CreateProductData) (*models.Product, error) {
	// ✅ Security: Verify the shop belongs to the user
	shop, err := s.shopRepo.FindByID(ctx, data.ShopID)
	if err != nil {
		return nil, err
	}

	if shop == nil {
		return nil, errors.New("shop not found")
	}

	if shop.UserID != userID {
		return nil, errors.New("unauthorized: shop does not belong to user")
	}

	// Create product
	product := &models.Product{
		ShopID:      data.ShopID,
		Product:     data.Product,
		Description: data.Description,
		Quantity:    data.Quantity,
	}

	err = s.repo.Create(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) Update(ctx context.Context, userID int64, data UpdateProductData) (*models.Product, error) {
	// ✅ Security: Verify the product exists
	product, err := s.repo.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if product == nil {
		return nil, errors.New("product not found")
	}

	// ✅ Security: Verify the shop belongs to the user
	shop, err := s.shopRepo.FindByID(ctx, product.ShopID)
	if err != nil {
		return nil, err
	}

	if shop == nil {
		return nil, errors.New("shop not found")
	}

	if shop.UserID != userID {
		return nil, errors.New("unauthorized: shop does not belong to user")
	}

	// Update product
	product.Product = data.Product
	product.Description = data.Description
	product.Quantity = data.Quantity

	err = s.repo.Update(ctx, product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

func (s *service) Delete(ctx context.Context, userID int64, id int64) error {
	// ✅ Security: Verify the product exists
	product, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if product == nil {
		return errors.New("product not found")
	}

	// ✅ Security: Verify the shop belongs to the user
	shop, err := s.shopRepo.FindByID(ctx, product.ShopID)
	if err != nil {
		return err
	}

	if shop == nil {
		return errors.New("shop not found")
	}

	if shop.UserID != userID {
		return errors.New("unauthorized: shop does not belong to user")
	}

	return s.repo.Delete(ctx, product)
}