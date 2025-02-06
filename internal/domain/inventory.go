package domain

import (
	"database/sql"
	"time"
)

// Hub - Struct representing a hub
type Hub struct {
	ID        int64   `json:"id"` // Corresponds to BIGSERIAL (int64 in Go)
	Name      string  `json:"name"`
	TenantID  int64   `json:"tenant_id"`  // Corresponds to BIGINT (int64 in Go)
	Location  string  `json:"location"`   // Could be a string or use a custom type for the POINT type
CreatedAt *time.Time `json:"created_at,omitempty"`
	
	CreatedBy *time.Time  `json:"created_by"` // Corresponds to BIGINT (int64 in Go)
	UpdatedAt *time.Time   `json:"updated_at"` // Corresponds to TIMESTAMPTZ (string or time.Time in Go)
	UpdatedBy int64   `json:"updated_by"` // Corresponds to BIGINT (int64 in Go)
	DeletedAt *string `json:"deleted_at"` // Nullable, so it's a pointer to string or time.Time
}
type Sku struct {
	ID          int64     `json:"id"`
	SellerID    int64     `json:"seller_id"`
	Attributes  string    `json:"attributes"`
	PPU         float64   `json:"ppu"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}

type Inventory struct {
	ID        int64        `json:"id" db:"id"`
	HubID     int64        `json:"hub_id" db:"hub_id"`
	SKUID     int64        `json:"sku_id" db:"sku_id"`
	Quantity  int          `json:"quantity" db:"quantity"`
	CreatedAt sql.NullTime `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt sql.NullTime `json:"updated_at,omitempty" db:"updated_at"`
}