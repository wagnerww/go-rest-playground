package entity

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	id          string  `json:"email" gorm:"id"`
	description string  `json:"descricao"`
	preco       float64 `json:"preco"`
}
