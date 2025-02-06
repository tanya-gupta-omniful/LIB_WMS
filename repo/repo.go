package repo

import (
	"WMS/internal/domain"
	"context"
	"errors"
	"sync"

	"github.com/omniful/go_commons/db/sql/postgres"
)

type Repository interface {
	GetAllHubs(ctx context.Context) []domain.Hub
	GetAllSkus(ctx context.Context) []domain.Sku
	GetHubByID(ctx context.Context, id int) (domain.Hub, error) 
	GetSkuByID(ctx context.Context, skuID int) (domain.Sku, error)
	GetHubByTenantId(ctx context.Context, tenantId int)([]domain.Hub, error)
	GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.Sku, error)
	CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error)
	CreateSku(ctx context.Context, sku domain.Sku) (domain.Sku, error)
	DeleteHub(ctx context.Context, id int) error
	DeleteSku(ctx context.Context, skuID int) error
}

type repository struct {
	db *postgres.DbCluster
}

var repo *repository
var repoOnce sync.Once

// NewRepository is the constructor function that ensures the Repository is initialized only once.
func NewRepository(db *postgres.DbCluster) Repository {
	repoOnce.Do(func() {
		// Initialize the Repository with a given DbCluster.
		repo = &repository{
			db: db,
		}
	})
	return repo
}

func (r *repository) GetAllHubs(ctx context.Context) []domain.Hub {

	var hubs []domain.Hub
	//err := r.db.Find(&hubs).Error
	r.db.GetMasterDB(ctx).Find(&hubs)
	return hubs
}
func (r *repository)GetAllSkus(ctx context.Context)[]domain.Sku{
	var skus []domain.Sku
	r.db.GetMasterDB(ctx).Find(&skus)
	return skus
}
func (r *repository) GetHubByID(ctx context.Context, id int) (domain.Hub, error) {
	var hub domain.Hub
	// Check if the ID is valid (optional but useful to prevent unnecessary DB queries)
	if id <= 0 {
		return hub, errors.New("invalid ID")
	}
	
	// Query the database to fetch the hub with the given ID
	result := r.db.GetMasterDB(ctx).First(&hub, id)
	if result.Error != nil {
		// Return an empty Hub and the error if not found
		return domain.Hub{}, result.Error
	}
	
	return hub, nil
}

// GetSkuByID queries the database to fetch SKU by its ID
func (r *repository) GetSkuByID(ctx context.Context, skuID int) (domain.Sku, error) {
	var sku domain.Sku

	// Check if the SKU ID is valid
	if skuID <= 0 {
		return sku, errors.New("invalid SKU ID")
	}

	// Query the database using SKU ID
	result := r.db.GetMasterDB(ctx).Where("id = ?", skuID).First(&sku)
	if result.Error != nil {
		// Handle error if SKU is not found or any database error occurs
		return domain.Sku{}, result.Error
	}

	// Return the found SKU
	return sku, nil
}
func (r *repository) GetHubByTenantId(ctx context.Context,tenantId int) ([]domain.Hub, error){
	var hubs []domain.Hub
	if tenantId<=0 {
		return hubs, errors.New("invalid tenant ID")
	}
	result := r.db.GetMasterDB(ctx).Where("tenant_id = ?", tenantId).Find(&hubs)
	if result.Error != nil{
		return hubs, result.Error
	}	
	return hubs, nil
} 
// GetSkuBySellerID queries the database to fetch SKUs by seller_id
func (r *repository) GetSkuBySellerID(ctx context.Context, sellerID int) ([]domain.Sku, error) {
	var skus []domain.Sku

	// Check if the seller ID is valid
	if sellerID <= 0 {
		return skus, errors.New("invalid Seller ID")
	}

	// Query the database using seller_id to get multiple SKUs
	result := r.db.GetMasterDB(ctx).Where("seller_id = ?", sellerID).Find(&skus)
	if result.Error != nil {
		// Handle error if no SKUs are found for the given seller_id
		return nil, result.Error
	}

	// Return the found SKUs
	return skus, nil
}

// Create a new hub
func (r *repository) CreateHub(ctx context.Context, hub domain.Hub) (domain.Hub, error) {
	result := r.db.GetMasterDB(ctx).Create(&hub)
	return hub, result.Error
}

// Create a new SKU
func (r *repository) CreateSku(ctx context.Context, sku domain.Sku) (domain.Sku, error) {
	result := r.db.GetMasterDB(ctx).Create(&sku)
	return sku, result.Error
}

// Delete a hub by ID
func (r *repository) DeleteHub(ctx context.Context, id int) error {
	result := r.db.GetMasterDB(ctx).Delete(&domain.Hub{}, id)
	return result.Error
}

// Delete an SKU by ID
func (r *repository) DeleteSku(ctx context.Context, skuID int) error {
	result := r.db.GetMasterDB(ctx).Delete(&domain.Sku{}, skuID)
	return result.Error
}