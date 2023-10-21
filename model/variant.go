package model

import "time"

type Variant struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	VariantName string `json:"variant_name"`
	Quantity   uint     `json:"quantity"`
	ProductID  uint     `json:"product_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}