package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/wagnerww/go-rest-playground.git/models/dto/input"
	"github.com/wagnerww/go-rest-playground.git/service"
)

type ProductController interface {
	Save(c *gin.Context)
	GetAll(c *gin.Context)
}

type ProductControllerImpl struct {
	productService service.ProductService
}

func NewUProductController(productService service.ProductService) ProductController {
	return ProductControllerImpl{
		productService: productService,
	}
}

func (p ProductControllerImpl) Save(c *gin.Context) {
	var input input.Product
	err := c.BindJSON(&input)

	if err != nil {
		log.Fatalln(err)
		c.Status(http.StatusBadRequest)
	}

	c.JSON(http.StatusOK, p.productService.Save(input))
}

func (p ProductControllerImpl) GetAll(c *gin.Context) {
	log.Println("chegou controler")
	c.JSON(http.StatusOK, p.productService.GetAll())
}
