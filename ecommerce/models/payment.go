package models

import (
	"time"
)

type (
	// Payment
	Payment struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"provider"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// one to many
		SaleOrders []SaleOrder `json:"-"`
	}
)
