package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/app/service"
	"github.com/meziaris/devstore/internal/pkg/handler"
)

type ProductController struct {
	service service.IProductService
}

func NewProductController(service service.IProductService) *ProductController {
	return &ProductController{service: service}
}

// create new product
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	req := &schema.CreateProductReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	imageULR, err := c.service.Create(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.CreateProductResp{
		Name:        req.Name,
		Description: req.Description,
		Currency:    req.Currency,
		TotalStock:  req.TotalStock,
		IsActive:    req.IsActive,
		CategoryID:  req.CategoryID,
		Image:       imageULR,
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success ceate product", resp)
}

// get all Product
func (cc *ProductController) BrowseProduct(ctx *gin.Context) {
	req := schema.BrowseProductReq{
		Page:  ctx.GetInt("page"),
		Limit: ctx.GetInt("limit"),
	}
	resp, err := cc.service.BrowseAll(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get product list", resp)
}

// get detail Product
func (cc *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success found the product", resp)
}

// update product by id
func (cc *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateProductReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	imageURL, err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	resp := schema.UpdateProductResp{
		Name:        req.Name,
		Description: req.Description,
		Currency:    req.Currency,
		TotalStock:  req.TotalStock,
		IsActive:    req.IsActive,
		CategoryID:  req.CategoryID,
		Image:       imageURL,
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", resp)
}

// delete product by id
func (cc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete product", nil)
}
