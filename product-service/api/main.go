package main

import (
	"github.com/wagnerww/go-rest-playground.git/api/routes"
	"github.com/wagnerww/go-rest-playground.git/infraestructure/database"
)

func main() {
	db := database.Init()

	routes.SetupRoutes(db)

}
