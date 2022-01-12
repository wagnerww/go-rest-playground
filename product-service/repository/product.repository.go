package repository

import (
	"github.com/wagnerww/go-rest-playground.git/models/entity"
	"github.com/wagnerww/go-rest-playground.git/repository/impl"
	"gorm.io/gorm"
)

type ProductRepository interface {
	Save(entity.Product) (entity.Product, error)
	GetAll() []entity.Product
}

func NewProductRepository(DB *gorm.DB) ProductRepository {
	return impl.ProductRepositoryImpl{DB}
}
