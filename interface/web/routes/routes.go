package routes

import (
	"C"

	"github.com/gin-gonic/gin"
	"github.com/koutsoumposval/ddd-playground-golang/interface/web/controller/product"
)

// Router returns a http router with endpoints set
func Router() *gin.Engine {

	router := gin.Default()

	router.GET("/product/:id", product.GetProduct)
	router.GET("/product", product.GetProducts)
	router.POST("/product", product.AddProduct)

	return router
}
