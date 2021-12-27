package models

import (
	"time"
)

type (
	// UnitOfMeasure
	UnitOfMeasure struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"nama_unit"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// one to many
		Products []Product `json:"-"`
	}
)
