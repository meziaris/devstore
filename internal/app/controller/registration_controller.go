package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/pkg/handler"
)

type Registerer interface {
	Register(req *schema.RegisterReq) error
}

type RegistrationController struct {
	service Registerer
}

func NewRegistrationController(service Registerer) *RegistrationController {
	return &RegistrationController{service: service}
}

func (c *RegistrationController) Register(ctx *gin.Context) {
	req := schema.RegisterReq{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	if err := c.service.Register(&req); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success register", req)
}
