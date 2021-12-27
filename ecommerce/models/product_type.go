package models

import (
	"time"
)

type (
	// ProductType
	ProductType struct {
		ID        uint      `gorm:"primary_key" json:"id"`
		Name      string    `json:"type"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// one to many
		Products []Product `json:"-"`
	}
)
