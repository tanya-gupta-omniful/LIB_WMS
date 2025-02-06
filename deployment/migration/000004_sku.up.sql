CREATE TABLE IF NOT EXISTS sku (
    id  BIGSERIAL PRIMARY KEY,
    seller_id BIGINT NOT NULL,
    attributes JSONB, 
    ppu DECIMAL NOT NULL,
    created_at TIMESTAMP DEFAULT NOW(),
    updated_at TIMESTAMP DEFAULT NOW(),
    FOREIGN KEY (seller_id) REFERENCES sellers(id) ON DELETE CASCADE
);

type Sku struct {
	ID          int64     `json:"id"`
	SellerID    int64     `json:"seller_id"`
	Attributes  string    `json:"attributes"`
	PPU         float64   `json:"ppu"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   *time.Time `json:"deleted_at"`
}