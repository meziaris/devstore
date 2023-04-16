package repository

import "github.com/meziaris/devstore/internal/app/model"

type ICategoryRepository interface {
	Create(category model.Category) error
	Browse() ([]model.Category, error)
	Update(category model.Category) error
	GetByID(id string) (model.Category, error)
	DeleteByID(id string) error
}

type IProductRepository interface {
	Create(product model.Product) error
	Browse() ([]model.Product, error)
	GetByID(id string) (model.Product, error)
	Update(product model.Product) error
	DeleteByID(id string) error
}
