package application

import (
	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
	"github.com/koutsoumposval/ddd-playground-golang/domain/repository"
	"github.com/koutsoumposval/ddd-playground-golang/domain/value"
)

// ProductAppSvc is the entrypoint for Product use cases
type ProductAppSvc struct {
	Repository repository.ProductRepository
}

var productAppSvc ProductAppSvc

// GetProductAppSvc returns ProductAppSvc
func GetProductAppSvc() ProductAppSvc {
	return productAppSvc
}

// GetProduct returns product
func (i ProductAppSvc) GetProduct(id value.ProductID) (*entity.Product, error) {

	return i.Repository.Get(id)
}

// GetProducts returns product list
func (i ProductAppSvc) GetProducts() ([]*entity.Product, error) {

	return i.Repository.GetAll()
}

// CreateProduct creates a new product
func (i ProductAppSvc) CreateProduct(name string, CategoryID value.CategoryID) error {

	p, err := entity.CreateProduct(name, CategoryID)

	if err != nil {
		return err
	}

	return i.Repository.Save(p)
}
