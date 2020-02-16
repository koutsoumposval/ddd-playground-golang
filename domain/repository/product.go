package repository

import (
	"github.com/koutsoumposval/polyglot-ddd-product/domain/entity"
	"github.com/koutsoumposval/polyglot-ddd-product/domain/value"
)

// ProductRepository represent a repository of product
type ProductRepository interface {
	Get(id value.ProductID) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Save(product *entity.Product) error
}
