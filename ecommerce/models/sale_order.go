package models

import (
	"time"
)

//  membuat model, untuk manipulasi data. representasi datanya
type (
	// SaleOrder
	SaleOrder struct {
		ID              uint      `json:"id" gorm:"primary_key"`
		Code            string    `json:"kode_pesanan"`
		Status          string    `json:"status"`
		OrderDate       time.Time `json:"order_date"`
		ShippingAddress string    `json:"shipping_address"`
		CreatedAt       time.Time `json:"created_at"`
		UpdatedAt       time.Time `json:"updated_at"`
		// FK
		UserID    uint `json:"users_id"`
		PaymentID uint `json:"payment_id"`
		// many to one
		User    User    `json:"-"`
		Payment Payment `json:"-"`
		// one to many
		SaleOrderLines []SaleOrderLine `json:"-"`
	}
)
