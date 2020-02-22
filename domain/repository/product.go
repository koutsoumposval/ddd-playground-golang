package repository

import (
	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
	"github.com/koutsoumposval/ddd-playground-golang/domain/value"
)

// IProductRepository contract for product repository
type IProductRepository interface {
	Get(id value.ProductID) (*entity.Product, error)
	GetAll() ([]*entity.Product, error)
	Save(p *entity.Product) error
}
