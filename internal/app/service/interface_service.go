package service

import (
	"github.com/meziaris/devstore/internal/app/schema"
)

type ICategoryService interface {
	Create(req schema.CreateCategoryReq) error
	BrowseAll(req schema.BrowseCategoryReq) ([]schema.GetCategoryResp, error)
	GetByID(id string) (schema.GetCategoryResp, error)
	UpdateByID(id string, req schema.UpdateCategoryReq) error
	DeleteByID(id string) error
}

type IProductService interface {
	Create(req *schema.CreateProductReq) (imageURL string, err error)
	BrowseAll(req schema.BrowseProductReq) ([]schema.BrowseProductResp, error)
	GetByID(id string) (schema.DetailProductResp, error)
	UpdateByID(id string, req *schema.UpdateProductReq) (imageURL string, err error)
	DeleteByID(id string) error
}
