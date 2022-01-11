package repository

import (
	"log"

	"github.com/jinzhu/gorm"
	"github.com/wagnerww/go-rest-playground.git/models/entity"
)

type productRepository struct {
	DB *gorm.DB
}

type ProductRepository interface {
	Save(entity.Product) (entity.Product, error)
	GetAll() ([]entity.Product, error)
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	return productRepository{
		DB: db,
	}
}

func (p productRepository) Save(user entity.Product) (entity.Product, error) {
	log.Print("[productRepository]...Save")
	err := p.DB.Create(&user).Error
	return user, err

}

func (p productRepository) GetAll() (users []entity.Product, err error) {
	log.Print("[productRepository]...Get All")
	err = p.DB.Find(&users).Error
	return users, err
}
