package service

import (
	"WMS/internal/domain"
	"WMS/repo"
	"context"
	"errors"
)

type InventoryService interface {
	FetchInventory(ctx context.Context, hubID, skuID int) ([]domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventory domain.Inventory) error
	ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error)
}

type inventoryService struct {
	repo repo.InventoryRepository
}

// Constructor function
func NewInventoryService(r repo.InventoryRepository) InventoryService {
	return &inventoryService{
		repo: r,
	}
}

// Fetch inventory based on tenant_id, hub_id, and sku_id
func (s *inventoryService) FetchInventory(ctx context.Context, hubID, skuID int) ([]domain.Inventory, error) {
	return s.repo.FetchInventory(ctx, hubID, skuID)
}
func (s *inventoryService) UpdateInventory(ctx context.Context, inventory domain.Inventory) error {
	if  inventory.HubID == 0 || inventory.SKUID == 0 {
		return errors.New("invalid inventory data")
	}
	return s.repo.UpdateInventory(ctx, inventory)
}
func (s *inventoryService) ValidateInventory(ctx context.Context, skuID, hubID, quantity int) (bool, error) {
	return s.repo.ValidateInventory(ctx, skuID, hubID, quantity)
}