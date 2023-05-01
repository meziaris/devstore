package service

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/meziaris/devstore/internal/app/model"
	"github.com/meziaris/devstore/internal/app/repository"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/pkg/reason"
)

type ProductService struct {
	productRepo  repository.IProductRepository
	categoryRepo repository.ICategoryRepository
}

func NewProductService(pr repository.IProductRepository, cr repository.ICategoryRepository) *ProductService {
	return &ProductService{productRepo: pr, categoryRepo: cr}
}

func (s *ProductService) Create(req *schema.CreateProductReq) error {
	inserData := model.Product{
		Name:        req.Name,
		Description: req.Description,
		Currency:    req.Currency,
		TotalStock:  req.TotalStock,
		IsActive:    req.IsActive,
		CategoryID:  req.CategoryID,
	}

	categoryID := strconv.Itoa(req.CategoryID)
	if _, err := s.categoryRepo.GetByID(categoryID); err != nil {
		return errors.New(reason.CategoryNotFound)
	}

	if err := s.productRepo.Create(inserData); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) BrowseAll() ([]schema.BrowseProductResp, error) {
	var resp []schema.BrowseProductResp

	products, err := s.productRepo.Browse()
	if err != nil {
		return resp, errors.New(reason.ProductCannotBrowse)
	}

	for _, value := range products {
		respData := schema.BrowseProductResp{
			ID:          value.ID,
			Name:        value.Name,
			Description: value.Description,
			Currency:    value.Currency,
			TotalStock:  value.TotalStock,
			IsActive:    value.IsActive,
		}
		resp = append(resp, respData)
	}

	return resp, nil
}

func (s *ProductService) GetByID(id string) (schema.DetailProductResp, error) {
	var resp schema.DetailProductResp

	product, err := s.productRepo.GetByID(id)
	if err != nil {
		return resp, errors.New(reason.CategoryCannotGetDetail)
	}

	categoryID := strconv.Itoa(product.CategoryID)
	category, err := s.categoryRepo.GetByID(categoryID)
	if err != nil {
		return resp, errors.New(reason.CategoryNotFound)
	}

	resp = schema.DetailProductResp{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Currency:    product.Currency,
		TotalStock:  product.TotalStock,
		IsActive:    product.IsActive,
		Category: schema.Category{
			ID:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
	}

	return resp, nil
}

func (s *ProductService) UpdateByID(id string, req *schema.UpdateProductReq) error {
	updateData := model.Product{}

	oldData, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	updateData.ID = oldData.ID
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Currency = req.Currency
	updateData.TotalStock = req.TotalStock
	updateData.IsActive = req.IsActive
	updateData.CategoryID = req.CategoryID

	if err = s.productRepo.Update(updateData); err != nil {
		fmt.Println(err)
		return errors.New(reason.CategoryCannotUpdate)
	}

	return nil
}
func (s *ProductService) DeleteByID(id string) error {

	_, err := s.productRepo.GetByID(id)
	if err != nil {
		return errors.New(reason.ProductNotFound)
	}

	if err := s.productRepo.DeleteByID(id); err != nil {
		return errors.New(reason.ProductCannotDelete)
	}

	return nil
}
