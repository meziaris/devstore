package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/app/service"
	"github.com/meziaris/devstore/internal/pkg/handler"
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

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := cc.service.Create(req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success create category", req)
}

// get all category
func (cc *CategoryController) BrowseCategory(ctx *gin.Context) {
	resp, err := cc.service.BrowseAll()
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot get category list")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": resp})
}

// get detail category
func (cc *CategoryController) DetailCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	resp, err := cc.service.GetByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, "cannot get category detail")
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success found the category", resp)
}

// update category by id
func (cc *CategoryController) UpdateCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")
	req := schema.UpdateCategoryReq{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	err := cc.service.UpdateByID(id, req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success update category", nil)
}

// delete category by id
func (cc *CategoryController) DeleteCategory(ctx *gin.Context) {
	id, _ := ctx.Params.Get("id")

	err := cc.service.DeleteByID(id)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success delete category", nil)
}
