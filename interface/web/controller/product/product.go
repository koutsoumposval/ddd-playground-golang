package product

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/koutsoumposval/ddd-playground-golang/application"
	"github.com/koutsoumposval/ddd-playground-golang/domain/value"
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

	pID := new(value.ProductID)
	pID.SetID(id)

	p, err := pc.ProductAppSvc.GetProduct(*pID)

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
		CategoryID int64  `json:"categoryId"`
	}{}

	err := json.Unmarshal(rawData, &data)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)

		return
	}

	cID := new(value.CategoryID)
	cID.SetID(data.CategoryID)

	err = pc.ProductAppSvc.CreateProduct(data.Name, *cID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, err)

		return
	}

	c.JSON(http.StatusOK, nil)
}
