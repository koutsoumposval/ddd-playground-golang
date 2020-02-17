package routes

import (
	"C"

	"github.com/gin-gonic/gin"
	"github.com/koutsoumposval/ddd-playground-golang/interface/web/controller/product"
)

// Router returns a http router with endpoints set
func Router() *gin.Engine {

	r := gin.Default()

	r.GET("/product/:id", product.GetProduct)
	r.GET("/product", product.GetProducts)
	r.POST("/product", product.AddProduct)

	return r
}
