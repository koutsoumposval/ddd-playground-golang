package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/koutsoumposval/ddd-playground-golang/application"
)

// ProductController handles requests and invokes the specific use case
type ProductController struct {
	ProductAppSvc application.IProductAppSvc
}

// GetProduct gets a product
func (pc ProductController) GetProduct(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	p, err := pc.ProductAppSvc.GetProduct(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, p)
}

// GetProducts gets all products
func (pc ProductController) GetProducts(c *gin.Context) {

	ps, err := pc.ProductAppSvc.GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ps)
}

// AddProduct adds a product
func (pc ProductController) AddProduct(c *gin.Context) {

	rawData, _ := c.GetRawData()

	data := struct {
		Name       string `json:"name"`
		CategoryID int64  `json:"category_id"`
	}{}

	err := json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	err = pc.ProductAppSvc.CreateProduct(data.Name, data.CategoryID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, nil)
}
