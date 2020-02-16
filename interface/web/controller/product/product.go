package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/koutsoumposval/polyglot-ddd-product/application"
	"github.com/koutsoumposval/polyglot-ddd-product/domain/value"
)

// GetProduct gets a product
func GetProduct(c *gin.Context) {

	id, err := strconv.ParseInt(c.Param("id"), 10, 64)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	pID := new(value.ProductID)
	pID.SetID(id)

	p, err := application.GetProductAppSvc().GetProduct(*pID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, p)
}

// GetProducts gets all products
func GetProducts(c *gin.Context) {

	ps, err := application.GetProductAppSvc().GetProducts()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, ps)
}

// AddProduct adds a product
func AddProduct(c *gin.Context) {

	rawData, _ := c.GetRawData()

	data := struct {
		Name       string `json:"name"`
		CategoryID int64  `json:"categoryId"`
	}{}

	err := json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	cID := new(value.CategoryID)
	cID.SetID(data.CategoryID)

	err = application.GetProductAppSvc().CreateProduct(data.Name, *cID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, nil)
}
