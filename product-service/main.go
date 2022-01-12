package main

import (
	"github.com/wagnerww/go-rest-playground.git/infraestructure/database"
	"github.com/wagnerww/go-rest-playground.git/routes"
)

func main() {
	db := database.Init()

	routes.SetupRoutes(db)

}
