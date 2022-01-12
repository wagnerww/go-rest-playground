package impl

import (
	"log"

	"github.com/wagnerww/go-rest-playground.git/models/entity"
	"gorm.io/gorm"
)

type ProductRepositoryImpl struct {
	DB *gorm.DB
}

func (p ProductRepositoryImpl) Save(product entity.Product) (entity.Product, error) {
	log.Print("[productRepository]...Save")
	err := p.DB.Create(&product).Error
	return product, err
}

func (p ProductRepositoryImpl) GetAll() []entity.Product {
	log.Print("[productRepository]...Get All")
	var products []entity.Product
	p.DB.Find(&products)
	return products
}
