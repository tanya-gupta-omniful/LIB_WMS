package domain

// Hub - Struct representing a hub
type Hub struct {
	ID        int64   `json:"id"` // Corresponds to BIGSERIAL (int64 in Go)
	Name      string  `json:"name"`
	TenantID  int64   `json:"tenant_id"`  // Corresponds to BIGINT (int64 in Go)
	Location  string  `json:"location"`   // Could be a string or use a custom type for the POINT type
	CreatedAt string  `json:"created_at"` // Corresponds to TIMESTAMPTZ (string or time.Time in Go)
	CreatedBy int64   `json:"created_by"` // Corresponds to BIGINT (int64 in Go)
	UpdatedAt string  `json:"updated_at"` // Corresponds to TIMESTAMPTZ (string or time.Time in Go)
	UpdatedBy int64   `json:"updated_by"` // Corresponds to BIGINT (int64 in Go)
	DeletedAt *string `json:"deleted_at"` // Nullable, so it's a pointer to string or time.Time
}