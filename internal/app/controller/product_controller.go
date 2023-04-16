package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/controller/internal"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/app/service"
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
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	if err := c.service.Create(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, internal.APIResponse(http.StatusOK, "success create product", req))
}

// get all Product
func (cc *ProductController) BrowseProduct(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// get detail Product
func (cc *ProductController) DetailProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// update article by id
func (cc *ProductController) UpdateProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := &schema.UpdateProductReq{}

	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	err = cc.service.UpdateByID(id, req)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, internal.APIResponse(http.StatusOK, "success update product", req))
}

// delete article by id
func (cc *ProductController) DeleteProduct(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, internal.APIResponse(http.StatusOK, "success delete product"))
}
