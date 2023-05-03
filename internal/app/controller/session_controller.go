package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/pkg/handler"
	"github.com/meziaris/devstore/internal/pkg/reason"
)

type SessionService interface {
	Login(req *schema.LoginReq) (schema.LoginResp, error)
	Logout(userID int) error
	Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error)
}

type RefreshTokenVerifier interface {
	VerifyRefreshToken(tokenString string) (string, error)
}

type SessionController struct {
	sessionService SessionService
	tokenCreator   RefreshTokenVerifier
}

func NewSessionController(sessionService SessionService, tokenCreator RefreshTokenVerifier) *SessionController {
	return &SessionController{sessionService: sessionService, tokenCreator: tokenCreator}
}

func (c *SessionController) Login(ctx *gin.Context) {
	req := schema.LoginReq{}

	if handler.BindAndCheck(ctx, &req) {
		return
	}

	res, err := c.sessionService.Login(&req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success login", res)
}

// refresh
func (c *SessionController) Refresh(ctx *gin.Context) {
	refreshToken := ctx.GetHeader("refresh_token")
	if refreshToken == "" {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
	}

	sub, err := c.tokenCreator.VerifyRefreshToken(refreshToken)

	if err != nil {
		handler.ResponseError(ctx, http.StatusUnauthorized, reason.FailedRefreshToken)
		return
	}

	intSub, _ := strconv.Atoi(sub)
	req := &schema.RefreshTokenReq{}
	req.RefreshToken = refreshToken
	req.UserID = intSub

	res, err := c.sessionService.Refresh(req)
	if err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, reason.FailedRefreshToken)
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success refresh", res)
}

func (c *SessionController) Logout(ctx *gin.Context) {
	userID, _ := strconv.Atoi(ctx.GetString("user_id"))
	if err := c.sessionService.Logout(userID); err != nil {
		handler.ResponseError(ctx, http.StatusUnprocessableEntity, err.Error())
		return
	}

	handler.ResponseSuccess(ctx, http.StatusOK, "success logout", nil)
}
