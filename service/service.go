package service

import (
	"WMS/internal/domain"
	"WMS/repo"
	"context"
	"errors"
)

type Service interface {
	FetchHubs(ctx context.Context) []domain.Hub
	FetchHubByID(ctx context.Context, id int) (domain.Hub, error)
	FetchSkuByID(ctx context.Context, skuID int) (domain.Sku, error)
	FetchSkuBySellerID(ctx context.Context, sellerID int) ([]domain.Sku, error)
}

type service struct {
	repo repo.Repository
}

// NewService is the constructor function to create a new instance of ConcreteService.
func NewService(r repo.Repository) Service {
	return &service{
		repo: r,
	}
}

func (s *service) FetchHubs(ctx context.Context) []domain.Hub {
	return s.repo.GetAllHubs(ctx)
}

func (s *service) FetchHubByID(ctx context.Context, id int) (domain.Hub, error) { // Implementing the new method
	hub, err := s.repo.GetHubByID(ctx, id)
	if err != nil {
		return domain.Hub{}, err // Return an empty Hub and the error
	}
	return hub, nil
}

func (s *service) FetchSkuByID(ctx context.Context, skuID int) (domain.Sku, error) {
	// Check if skuID is valid
	if skuID <= 0 {
		return domain.Sku{}, errors.New("invalid SKU ID")
	}

	// Call repository to fetch SKU by SKU ID
	sku, err := s.repo.GetSkuByID(ctx, skuID)
	if err != nil {
		return domain.Sku{}, err // Return an empty SKU and the error
	}
	return sku, nil
}

// FetchSkuBySellerID fetches SKUs based on the Seller ID
func (s *service) FetchSkuBySellerID(ctx context.Context, sellerID int) ([]domain.Sku, error) {
	// Check if sellerID is valid
	if sellerID <= 0 {
		return nil, errors.New("invalid Seller ID")
	}

	// Call repository to fetch SKUs by Seller ID
	skus, err := s.repo.GetSkuBySellerID(ctx, sellerID)
	if err != nil {
		return nil, err // Return nil and the error
	}
	return skus, nil
}