package main

import (
	"fmt"

	"github.com/koutsoumposval/ddd-playground-golang/app/config"
	"github.com/koutsoumposval/ddd-playground-golang/application"
	"github.com/koutsoumposval/ddd-playground-golang/infrastructure/persistence"
	"github.com/koutsoumposval/ddd-playground-golang/interface/http/controller/product"
	"github.com/koutsoumposval/ddd-playground-golang/interface/http/routes"
)

func main() {

	conn, err := config.DatabaseConnection()

	if err != nil {
		fmt.Println("Error connecting to Database")
	}

	pRepositoryImpl := persistence.ProductRepository{
		Connection: conn,
	}

	pAppSvcImpl := application.ProductAppSvc{
		Repository: &pRepositoryImpl,
	}

	pControllerImpl := product.ProductController{
		ProductAppSvc: &pAppSvcImpl,
	}

	routes.Router(pControllerImpl).Run()
}
