package service

import (
	"errors"
	"mime/multipart"
	"strconv"

	"github.com/meziaris/devstore/internal/app/model"
	"github.com/meziaris/devstore/internal/app/repository"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/pkg/reason"
)

type ImageUploader interface {
	UploadImage(userID string, input *multipart.FileHeader) (imageURL string, err error)
}

type ProductService struct {
	productRepo   repository.IProductRepository
	categoryRepo  repository.ICategoryRepository
	imageUploader ImageUploader
}

func NewProductService(pr repository.IProductRepository, cr repository.ICategoryRepository, uploader ImageUploader) *ProductService {
	return &ProductService{
		productRepo:   pr,
		categoryRepo:  cr,
		imageUploader: uploader,
	}
}

func (s *ProductService) Create(req *schema.CreateProductReq) (imageURL string, err error) {
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
		return "", errors.New(reason.CategoryNotFound)
	}

	productID, err := s.productRepo.Create(inserData)

	if err != nil {
		return "", errors.New(reason.CategoryCannotCreate)
	}

	imageURL, err = s.imageUploader.UploadImage(strconv.Itoa(productID), req.Image)
	if err != nil {
		return "", errors.New(reason.CategoryCannotCreate)
	}

	if err := s.productRepo.UpdateImageURL(productID, imageURL); err != nil {
		return "", errors.New(reason.CategoryCannotCreate)
	}

	return imageURL, nil
}

func (s *ProductService) BrowseAll(req schema.BrowseProductReq) ([]schema.BrowseProductResp, error) {
	var resp []schema.BrowseProductResp

	dbSearch := model.BrowseProduct{Page: req.Page, Limit: req.Limit}
	products, err := s.productRepo.Browse(dbSearch)
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
			ImageURL:    value.ImageURL,
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
		ImageURL: product.ImageURL,
	}

	return resp, nil
}

func (s *ProductService) UpdateByID(id string, req *schema.UpdateProductReq) (imageURL string, err error) {
	updateData := model.Product{}

	oldData, err := s.productRepo.GetByID(id)
	if err != nil {
		return "", errors.New(reason.ProductNotFound)
	}

	updateData.ID = oldData.ID
	updateData.Name = req.Name
	updateData.Description = req.Description
	updateData.Currency = req.Currency
	updateData.TotalStock = req.TotalStock
	updateData.IsActive = req.IsActive
	updateData.CategoryID = req.CategoryID

	if err = s.productRepo.Update(updateData); err != nil {
		return "", errors.New(reason.ProductCannotUpdate)
	}

	imageURL, err = s.imageUploader.UploadImage(strconv.Itoa(updateData.ID), req.Image)
	if err != nil {
		return "", errors.New(reason.ProductCannotUpdate)
	}

	if err := s.productRepo.UpdateImageURL(updateData.ID, imageURL); err != nil {
		return "", errors.New(reason.ProductCannotUpdate)
	}

	return imageURL, nil
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
