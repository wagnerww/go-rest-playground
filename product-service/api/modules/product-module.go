package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/wagnerww/go-rest-playground.git/api/controller"
	"github.com/wagnerww/go-rest-playground.git/repository"
	"github.com/wagnerww/go-rest-playground.git/service"
	"gorm.io/gorm"
)

func SetupProductModule(db *gorm.DB, httpRouter *gin.Engine) {
	// wirings - PRODUCTS
	authService := service.NewAuthService()
	productRepository := repository.NewProductRepository(db)
	productController := controller.NewUProductController(
		service.NewProductService(productRepository, authService),
	)

	products := httpRouter.Group("products")
	products.GET("/", productController.GetAll)
	products.POST("/", productController.Save)

}
