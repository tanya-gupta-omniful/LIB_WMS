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
	GetHubByID(ctx context.Context, id int) (domain.Hub, error) 
	GetHubByTenantID(ctx context.Context, id int)(domain.Hub, error)
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

func (r *repository) GetHubByTenantID(ctx context.Context, id int) (domain.Hub, error) {
	var hub domain.Hub
	// Check if the ID is valid (optional but useful to prevent unnecessary DB queries)
	if id <= 0 {
		return hub, errors.New("invalid Tenant ID")
	}
	
	// Query the database to fetch the hub with the given ID
	result := r.db.GetMasterDB(ctx).Where("tenant_id = ?", id).First(&hub)
	if result.Error != nil {
		// Return an empty Hub and the error if not found
		return domain.Hub{}, result.Error
	}
	
	return hub, nil
}
