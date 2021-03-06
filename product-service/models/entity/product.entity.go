package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID          int     `json:"id" gorm:"primaryKey"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}
