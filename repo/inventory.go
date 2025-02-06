package repo

import (
	"WMS/internal/domain"
	"context"
	"errors"

	"github.com/omniful/go_commons/db/sql/postgres"
)

type InventoryRepository interface {
	FetchInventory(ctx context.Context, hubID, skuID int) ([]domain.Inventory, error)
	UpdateInventory(ctx context.Context, inventory domain.Inventory) error
}

type inventoryRepository struct {
	db *postgres.DbCluster
}

func NewInventoryRepository(db *postgres.DbCluster) InventoryRepository {
	return &inventoryRepository{
		db: db,
	}
}

// Fetch inventory based on tenant_id, hub_id, and sku_id
func (r *inventoryRepository) FetchInventory(ctx context.Context,  hubID, skuID int) ([]domain.Inventory, error) {
	var inventories []domain.Inventory
	query := r.db.GetMasterDB(ctx)

	
	if hubID != 0 {
		query = query.Where("hub_id = ?", hubID)
	}
	if skuID != 0 {
		query = query.Where("sku_id = ?", skuID)
	}

	result := query.Find(&inventories)
	return inventories, result.Error
}

func (r *inventoryRepository) UpdateInventory(ctx context.Context, inventory domain.Inventory) error {
	result := r.db.GetMasterDB(ctx).Model(&inventory).
		Where(" hub_id = ? AND sku_id = ?", inventory.HubID, inventory.SKUID).
		Updates(inventory)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return errors.New("no inventory record found to update")
	}

	return nil
}