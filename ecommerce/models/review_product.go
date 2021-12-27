package models

import (
	"time"
)

//  membuat model, untuk manipulasi data. representasi datanya
type (
	// ReviewProduct
	ReviewProduct struct {
		ID        uint      `json:"id" gorm:"primary_key"`
		Rating    int       `json:"rating"`
		Deskripsi string    `json:"deskripsi"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		// FK
		UserID    uint `json:"users_id"`
		ProductID uint `json:"product_id"`
		// many to one
		User    User    `json:"-"`
		Product Product `json:"-"`
	}
)
