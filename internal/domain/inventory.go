package domain

import (
	"context"
)

// Inventory - Struct representing inventory details
type Inventory struct {
	ID       uint64 `json:"id"`
	SKUCode  string `json:"sku_code"`
	HubID    uint64 `json:"hub_id"`
	Quantity int    `json:"quantity"`
}

// InventoryService - Interface defining inventory operations
type InventoryService interface {
	GetInventory(ctx context.Context, sellerID uint64, hubID uint64) ([]Inventory, error)
	UpdateInventory(ctx context.Context, inventory Inventory) error
}
