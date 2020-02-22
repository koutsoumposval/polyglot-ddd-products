package application

import (
	"github.com/koutsoumposval/ddd-playground-golang/domain/entity"
	"github.com/koutsoumposval/ddd-playground-golang/domain/repository"
	"github.com/koutsoumposval/ddd-playground-golang/domain/value"
)

// IProductAppSvc is entry point for product use cases
type IProductAppSvc interface {
	GetProduct(id value.ProductID) (*entity.Product, error)
	GetProducts() ([]*entity.Product, error)
	CreateProduct(name string, CategoryID value.CategoryID) error
}

// ProductAppSvc is implementation of IProductAppSvc
type ProductAppSvc struct {
	Repository repository.IProductRepository
}

// GetProductAppSvc returns initialized ProductAppSvc
func GetProductAppSvc(repository repository.IProductRepository) IProductAppSvc {
	return &ProductAppSvc{Repository: repository}
}

// GetProduct returns product
func (p ProductAppSvc) GetProduct(id value.ProductID) (*entity.Product, error) {

	return p.Repository.Get(id)
}

// GetProducts returns product list
func (p ProductAppSvc) GetProducts() ([]*entity.Product, error) {

	return p.Repository.GetAll()
}

// CreateProduct creates a new product
func (p ProductAppSvc) CreateProduct(name string, CategoryID value.CategoryID) error {

	pr, err := entity.CreateProduct(name, CategoryID)

	if err != nil {
		return err
	}

	return p.Repository.Save(pr)
}
