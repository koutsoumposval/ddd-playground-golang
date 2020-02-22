package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/koutsoumposval/ddd-playground-golang/interface/web/controller/product"
)

// Router returns a http router with endpoints set
func Router(pController product.ProductController) *gin.Engine {


	r := gin.Default()

	r.GET("/product/:id", pController.GetProduct)
	r.GET("/product", pController.GetProducts)
	r.POST("/product", pController.AddProduct)

	return r
}