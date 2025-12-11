package service

import (
	"Shop/models"
	"Shop/module/shop/repository"
	"context"
	"errors"
)

type service struct {
	repo repository.ShopRepository
}

func NewShopService(repo repository.ShopRepository) ShopService {
	return &service{repo: repo}
}

func (s *service) FindByID(ctx context.Context, id int64) (*models.Shop, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *service) FindByUserID(ctx context.Context, userID int64) ([]models.Shop, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *service) Create(ctx context.Context, userID int64, data CreateShopData) (*models.Shop, error) {
	// Create new shop (users can have multiple shops)
	shop := &models.Shop{
		UserID:  userID,
		Name:    data.Name,
		Address: data.Address,
	}

	err := s.repo.Create(ctx, shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *service) Update(ctx context.Context, userID int64, data UpdateShopData) (*models.Shop, error) {
	// ✅ Security: Verify the shop belongs to the user
	shop, err := s.repo.FindByID(ctx, data.ID)
	if err != nil {
		return nil, err
	}

	if shop == nil {
		return nil, errors.New("shop not found")
	}

	if shop.UserID != userID {
		return nil, errors.New("unauthorized: shop does not belong to user")
	}

	// Update shop
	shop.Name = data.Name
	shop.Address = data.Address

	err = s.repo.Update(ctx, shop)
	if err != nil {
		return nil, err
	}

	return shop, nil
}

func (s *service) Delete(ctx context.Context, userID int64, id int64) error {
	// ✅ Security: Verify the shop belongs to the user
	shop, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return err
	}

	if shop == nil {
		return errors.New("shop not found")
	}

	if shop.UserID != userID {
		return errors.New("unauthorized: shop does not belong to user")
	}

	return s.repo.Delete(ctx, shop)
}