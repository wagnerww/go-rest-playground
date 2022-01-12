package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/wagnerww/go-rest-playground.git/modules"
	"gorm.io/gorm"
)

func SetupRoutes(db *gorm.DB) {

	httpRouter := gin.Default()

	modules.SetupProductModule(db, httpRouter)

	err := httpRouter.Run()
	if err != nil {
		panic(err)
	}

}
