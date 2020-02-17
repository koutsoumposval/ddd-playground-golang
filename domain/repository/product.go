package repository

import (
	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
	"github.com/koutsoumposval/ddd-playground-golang/domain/value"
)

// ProductRepository represent a repository of product
type ProductRepository interface {
	Get(id value.ProductID) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Save(product *entity.Product) error
}
