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

	if err := c.service.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot create product")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success ceate product", req)
}

// get all Product
func (cc *ProductController) BrowseProduct(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot get product list")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get product list", resp)
}

// get detail Product
func (cc *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot get product detail")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success get product detail", resp)
}

// update article by id
func (cc *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateProductReq{}

	if handler.BindAndCheck(ctx, req) {
		return
	}

	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot update product")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update product", req)
}

// delete article by id
func (cc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot delete product")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete product", nil)
}
