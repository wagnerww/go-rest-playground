package service

import (
	"fmt"
	"log"

	"github.com/wagnerww/go-rest-playground.git/models/dto/input"
	"github.com/wagnerww/go-rest-playground.git/models/dto/output"
	"github.com/wagnerww/go-rest-playground.git/models/entity"
	"github.com/wagnerww/go-rest-playground.git/models/mapper"
	"github.com/wagnerww/go-rest-playground.git/repository"
)

type ProductService interface {
	Save(input.Product) output.Product
	GetAll() []output.Product
}

type ProductServiceImpl struct {
	productRepository repository.ProductRepository
}

func NewProductService(repository repository.ProductRepository) ProductService {
	return ProductServiceImpl{repository}
}

func (p ProductServiceImpl) Save(productInput input.Product) output.Product {
	var productEntity entity.Product
	var output output.Product
	productEntity = entity.Product{
		Description: productInput.Description,
		Price:       productInput.Price,
	}

	productEntity, err := p.productRepository.Save(productEntity)
	output = mapper.ToSimpleOuput(productEntity)
	fmt.Println("error", err)
	return output
}

func (p ProductServiceImpl) GetAll() []output.Product {
	log.Println("chegou service")
	var entities []entity.Product
	var outputs []output.Product
	entities = p.productRepository.GetAll()

	for i := range entities {
		var outputEntity = mapper.ToSimpleOuput(entities[i])
		outputs = append(outputs, outputEntity)
	}

	return outputs

}
