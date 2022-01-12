package mapper

import (
	"github.com/wagnerww/go-rest-playground.git/models/dto/output"
	"github.com/wagnerww/go-rest-playground.git/models/entity"
)

type ProductMapper interface {
	toSimpleOuput(entity.Product) entity.Product
}

func ToSimpleOuput(entity entity.Product) output.Product {

	return output.Product{
		ID:          entity.ID,
		Description: entity.Description,
		Price:       entity.Price,
	}

}
