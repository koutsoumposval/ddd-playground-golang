package entity

import (
	"fmt"

	"github.com/koutsoumposval/polyglot-ddd-product/domain/value"
)

// Product represent entity of a product
type Product struct {
	ID         value.ProductID  `json:"id"`
	Name       string           `json:"name"`
	CategoryID value.CategoryID `json:"category_id"`
}

// CreateProduct creates a product
func CreateProduct(name string, categoryID value.CategoryID) (*Product, error) {

	if name == "" {
		return nil, fmt.Errorf("Invalid product name provided")
	}

	if categoryID.ID.ID() <= 0 {
		return nil, fmt.Errorf("Invalid category id provided")
	}

	return &Product{
		Name:       name,
		CategoryID: categoryID,
	}, nil
}
