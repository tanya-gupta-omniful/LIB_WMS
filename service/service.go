package service

import (
	"WMS/internal/domain"
	"WMS/repo"
	"context"
	"errors"
)

type Service interface {
	FetchHubs(ctx context.Context) []domain.Hub
	FetchSkus(ctx context.Context) []domain.Sku
	FetchHubByID(ctx context.Context, id int) (domain.Hub, error)
	FetchSkuByID(ctx context.Context, skuID int) (domain.Sku, error)
	FetchHubByTenantId(ctx context.Context, tenantId int) ([]domain.Hub,error)
	FetchSkuBySellerID(ctx context.Context, sellerID int) ([]domain.Sku, error)
	CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error)
	CreateSku(ctx context.Context, sku domain.Sku) (domain.Sku, error)
	DeleteHub(ctx context.Context, id int) error
	DeleteSku(ctx context.Context, skuID int) error
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
func (s *service)FetchSkus(ctx context.Context) []domain.Sku{
	return s.repo.GetAllSkus(ctx)
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
func(s *service) FetchHubByTenantId(ctx context.Context, tenantId int)([]domain.Hub, error){
	return s.repo.GetHubByTenantId(ctx, tenantId)
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
func (s *service) CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) {
	return s.repo.CreateHub(ctx, hub)
}

func (s *service) CreateSku(ctx context.Context, sku domain.Sku) (domain.Sku, error) {
	return s.repo.CreateSku(ctx, sku)
}

func (s *service) DeleteHub(ctx context.Context, id int) error {
	return s.repo.DeleteHub(ctx, id)
}

func (s *service) DeleteSku(ctx context.Context, skuID int) error {
	return s.repo.DeleteSku(ctx, skuID)
}