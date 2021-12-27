package models

import (
	"time"
)

//  membuat model, untuk manipulasi data. representasi datanya
type (
	// Product
	Product struct {
		ID            uint      `json:"id" gorm:"primary_key"`
		Name          string    `json:"nama"`
		CostPrice     int       `json:"cost_price"`
		QtyAvailable  int       `json:"qty_available"`
		QtyForecasted int       `json:"qty_forecasted"`
		ImageUrl      string    `json:"image_url"`
		Keterangan    string    `json:"keterangan"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"updated_at"`
		// foreign key
		UnitOfMeasureID uint `json:"uom_id"`
		ProductTypeID   uint `json:"product_type_id"`
		// many to one
		UnitOfMeasure UnitOfMeasure `json:"-"`
		ProductType   ProductType   `json:"-"`
		// one to many
		ReviewProducts []ReviewProduct `json:"-"`
		SaleOrderLines []SaleOrderLine `json:"-"`
	}
)
