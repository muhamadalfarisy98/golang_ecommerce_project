package models

import (
	"time"
)

//  membuat model, untuk manipulasi data. representasi datanya
type (
	// SaleOrderLine
	SaleOrderLine struct {
		ID           uint      `json:"id" gorm:"primary_key"`
		JumlahBarang int       `json:"jumlah_barang"`
		JumlahHarga  int       `json:"jumlah_harga"`
		CreatedAt    time.Time `json:"created_at"`
		UpdatedAt    time.Time `json:"updated_at"`
		// FK
		ProductID   uint `json:"product_id"`
		SaleOrderID uint `json:"sale_order_id"`
		// many to one
		Product   Product   `json:"-"`
		SaleOrder SaleOrder `json:"-"`
	}
)
