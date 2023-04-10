package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/controller/internal"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/app/service"
)

type CategoryController struct {
	service service.ICategoryService
}

func NewCategoryController(service service.ICategoryService) *CategoryController {
	return &CategoryController{service: service}
}

// create category
func (cc *CategoryController) CreateCategory(ctx *gin.Context) {
	req := schema.CreateCategoryReq{}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	if err := cc.service.Create(req); err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, internal.APIResponse(http.StatusOK, "success create category", req))
}

// get all category
func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// get detail category
func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error(), nil))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// update article by id
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	var req schema.UpdateCategoryReq

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

	ctx.JSON(http.StatusCreated, internal.APIResponse(http.StatusOK, "success update category", req))
}

// delete article by id
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		ctx.JSON(http.StatusUnprocessableEntity, internal.APIResponse(http.StatusUnprocessableEntity, err.Error()))
		return
	}

	ctx.JSON(http.StatusCreated, internal.APIResponse(http.StatusOK, "success delete category"))
}
