package entity

import (
	"fmt"
)

// Product represent entity of a product
type Product struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	CategoryID int64  `json:"category_id"`
}

// CreateProduct creates a product
func CreateProduct(name string, categoryID int64) (*Product, error) {

	if name == "" {
		return nil, fmt.Errorf("Invalid product name provided")
	}

	if categoryID <= 0 {
		return nil, fmt.Errorf("Invalid category id provided")
	}

	return &Product{
		Name:       name,
		CategoryID: categoryID,
	}, nil
}
